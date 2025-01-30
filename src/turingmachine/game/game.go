package game

import (
	"errors"
	"io"
	"log/slog"
	"math"
	"math/bits"
	"math/rand/v2"
	"strconv"
	"strings"
)

const (
	// The maximum number of choices per game
	MaxNumberOfChoicesPerGame = 6
	// https://en.wikipedia.org/wiki/Base32#Crockford's_Base32
	base32encode = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"
	// When we generate user-facing unique IDs for a given game we want to make
	// it a bit harder for the user to visually recognize the solution to the
	// game. It's not meant to be impossible to do, just hard to do at a glance.
	encoderScramble uint64 = 0b010101101001111010111110110101011111100111100
	// a block of 5-bits
	block5 = 0b11111

	// The exponents to utilized for converting a game to a unique id
	gameExp1 uint64 = (MaxChoice + 1)
	gameExp2 uint64 = gameExp1 * gameExp1
	gameExp3 uint64 = gameExp2 * gameExp1
	gameExp4 uint64 = gameExp3 * gameExp1
	gameExp5 uint64 = gameExp4 * gameExp1
)

var (
	ErrGameEmpty                = errors.New("the game has no choices made")
	ErrGameChoiceOrderBlank     = errors.New("the game blank choices should be left at the end")
	ErrGameChoiceOrderCriterias = errors.New("the game criterias are not in ascending order")
	ErrGameRepatingCriteria     = errors.New("the game has a repeating criteria")
	ErrGameNoUniqueSolution     = errors.New("the game does not have a unique solution")
	ErrGameHasRedundant         = errors.New("the game has a redundant card")

	ErrGameStringLength = errors.New("a game string is always 9 bytes long")
)

var (
	// Decoder from string-ified game - populated on init.
	base32decode [math.MaxUint8 + 1]uint64
)

func init() {
	// Default decode
	for i, b := range base32encode {
		base32decode[b] = uint64(i)
	}
	// Decode lower-case
	for i, b := range strings.ToLower(base32encode) {
		base32decode[b] = uint64(i)
	}
	// Other non-zero special decodings from
	// https://en.wikipedia.org/wiki/Base32#Crockford's_Base32
	base32decode['i'] = 1
	base32decode['I'] = 1
	base32decode['l'] = 1
	base32decode['L'] = 1
}

// A game can be distilled to the choices of criteria+laws in order from the
// lowest to the highest.
// A game is always read from the lowest to the highest index of Choice and the
// first blank choice (if any) is the last considered.
type Game [MaxNumberOfChoicesPerGame]Choice

// Returns a game given a set of criteria cards and verification cards
func GameFromCards(criteriaCards []uint8, verificationCards []uint16) (game Game, ok bool) {
	// Check: the number of criteria cards is > 0
	// Check: the number of criteria cards is <= MaxNumberOfChoices
	// Check: there are the same number of criterias and verification cards
	n := len(criteriaCards)
	if n <= 0 || n > MaxNumberOfChoicesPerGame || n != len(verificationCards) {
		ok = false
		return
	}

	var i int
	for i = range n {
		game[i], ok = ChoiceFromCriteriaVerifier(criteriaCards[i], verificationCards[i])
		if !ok {
			break
		}
	}
	if ok {
		// Sort game in order to produce consistent IDs
		choices := 6
		for i = range MaxNumberOfChoicesPerGame {
			if game[i] == BlankChoice {
				choices--
			}
		}
		// Sort
		if choices == 6 {
			SortGame6(&game)
		} else if choices == 5 {
			SortGame5(&game)
		} else if choices == 4 {
			SortGame4(&game)
		} else if choices == 3 {
			SortGame3(&game)
		} else if choices == 2 {
			SortGame2(&game)
		}
	}
	return
}

// Generates a random solvable game with choices of a given difficulty.
// NOTE: choices MUST be in the range [4, 6] otherwise the function panics
func RandomSolvableGame(choices int, difficulty Difficulty) (game Game, err error) {
	maxChoice := byte(MaxChoice)
	if difficulty == EasyDifficulty {
		maxChoice = numberOfStandardChoice
	} else if difficulty == StandardDifficulty {
		maxChoice = numberOfStandardChoice
	}
	var u64 uint64
	for {
		u64 = rand.Uint64()
		game = Game{
			Choice(byte(u64)%maxChoice + 1),
			Choice(byte(u64>>8)%maxChoice + 1),
			Choice(byte(u64>>16)%maxChoice + 1),
			Choice(byte(u64>>24)%maxChoice + 1),
		}
		if choices >= 5 {
			game[4] = Choice(byte(u64>>32)%maxChoice + 1)
			if choices == 6 {
				game[5] = Choice(byte(u64>>40)%maxChoice + 1)
			}
		}
		// Check...
		// - that the number of unique criterias matches the number of
		// - if the game has the right difficulty
		// - if the game is solvable
		// - if the game contains redundant entries
		//
		// The checks are done in the order that performed best during testing
		if !game.HasUniqueSolution() || game.uniqueCriterias() != choices ||
			StateFromGame(game).HasRedundant() || game.Difficulty() != difficulty {
			continue
		}
		// Sort the game before returning it
		if choices == 6 {
			SortGame6(&game)
		} else if choices == 5 {
			SortGame5(&game)
		} else {
			SortGame4(&game)
		}
		return
	}
}

// Reads the (idx-1)th game from a reader
func GameFromReader(r io.ReaderAt, idx int64) (game Game, err error) {
	gameRaw := [MaxNumberOfChoicesPerGame]byte{}
	_, err = r.ReadAt(gameRaw[:], idx*MaxNumberOfChoicesPerGame)
	if err != nil {
		return
	}
	game[0] = Choice(gameRaw[0])
	game[1] = Choice(gameRaw[1])
	game[2] = Choice(gameRaw[2])
	game[3] = Choice(gameRaw[3])
	game[4] = Choice(gameRaw[4])
	game[5] = Choice(gameRaw[5])
	return
}

// Game value as int, useful for sorting.
// Requres 64-bit ints
func (game Game) Value() int {
	if strconv.IntSize != 64 {
		panic("this function can only be called on 64-bit builds")
	}
	return int(game[0]) | int(game[1])<<8 | int(game[2])<<16 |
		int(game[3])<<24 | int(game[4])<<32 | int(game[5])<<40
}

// Derives a Game from a game string.
// Note: the string MUST follow Crockford's Base32 alphabet and must be a valid
// game.
func GameFromString(gameStr string) (Game, error) {
	// Check length
	if len(gameStr) != 9 {
		return Game{}, ErrGameStringLength
	}

	// Decode uid
	var uid uint64 = base32decode[gameStr[0]] + (base32decode[gameStr[1]] << 5) +
		(base32decode[gameStr[2]] << 10) + (base32decode[gameStr[3]] << 15) +
		(base32decode[gameStr[4]] << 20) + (base32decode[gameStr[5]] << 25) +
		(base32decode[gameStr[6]] << 30) + (base32decode[gameStr[7]] << 35) +
		(base32decode[gameStr[8]] << 40)
	uid ^= encoderScramble

	// Convert back into a game
	return Game{
		Choice((uid / gameExp4) % gameExp1),
		Choice((uid / gameExp2) % gameExp1),
		Choice(uid % gameExp1),
		Choice((uid / gameExp5) % gameExp1),
		Choice((uid / gameExp1) % gameExp1),
		Choice((uid / gameExp3) % gameExp1),
	}, nil
}

// Sorts a game to make it more likely to pass strict validation
func (game *Game) Sort() {
	// Move all blank choices to the end
	start := 0
	end := MaxNumberOfChoicesPerGame
	for start < end {
		if game[start] == BlankChoice {
			end--
			game[start], game[end] = game[end], game[start]
			continue
		}
		start++
	}
	if end < MaxNumberOfChoicesPerGame && game[end] != BlankChoice {
		end++
	}
	// Sort the ramining choices
	if end == 6 {
		SortGame6(game)
	} else if end == 5 {
		SortGame5(game)
	} else if end == 4 {
		SortGame4(game)
	} else if end == 3 {
		SortGame3(game)
	} else if end == 2 {
		SortGame2(game)
	}
}

// Returns the string representation of a game.
// Note: the game MUST be valid.
func (game Game) String() string {
	// Generated a scrambled unique ID representing this game among all other
	// valid games.
	uid := uint64(game[2]) + uint64(game[4])*gameExp1 + uint64(game[1])*gameExp2 +
		uint64(game[5])*gameExp3 + uint64(game[0])*gameExp4 + uint64(game[3])*gameExp5
	uid ^= encoderScramble

	// Encode using the base32 encoder
	return string([]byte{
		base32encode[uid&block5],
		base32encode[(uid>>5)&block5],
		base32encode[(uid>>10)&block5],
		base32encode[(uid>>15)&block5],
		base32encode[(uid>>20)&block5],
		base32encode[(uid>>25)&block5],
		base32encode[(uid>>30)&block5],
		base32encode[(uid>>35)&block5],
		base32encode[(uid>>40)&block5],
	})
}

// Writes the game to a byte slice.
// Note: the byte slice MUST be of length MaxNumberOfChoices
func (game Game) WriteTo(s []byte, startingIdx int) {
	s[startingIdx+5] = byte(game[5]) // Forcing length check here
	s[startingIdx] = byte(game[0])
	s[startingIdx+1] = byte(game[1])
	s[startingIdx+2] = byte(game[2])
	s[startingIdx+3] = byte(game[3])
	s[startingIdx+4] = byte(game[4])
}

// Returns the number of unique criterias found in this game
func (game Game) uniqueCriterias() int {
	return bits.OnesCount64(game[0].CriteriaIdMask() | game[1].CriteriaIdMask() |
		game[2].CriteriaIdMask() | game[3].CriteriaIdMask() |
		game[4].CriteriaIdMask() | game[5].CriteriaIdMask())
}

// Returns a debug string
func (game Game) Debug() string {
	sb := strings.Builder{}
	for i := range MaxNumberOfChoicesPerGame {
		sb.WriteString(game[i].Debug())
	}
	return sb.String()
}

// Returns the number of choices up to the last non-blank choice.
func (game Game) NumberOfChoices() int {
	// Check in reverse order as we are more likely to be on the upper rang
	for i := MaxNumberOfChoicesPerGame - 1; i >= 0; i-- {
		if game[i] != BlankChoice {
			return i + 1
		}
	}
	return 0
}

// Returns true if all non-blank choices are valid.
// As a special case we'll validate all choices even those that should be
// ignored.
func (game Game) IsValid() bool {
	return game[0].IsValid() && game[1].IsValid() && game[2].IsValid() &&
		game[3].IsValid() && game[4].IsValid() && game[5].IsValid()
}

// Performs a strict validation and returns an error if anything is wrong
func (game Game) ValidateStrict() error {
	// At least one choice
	if game[0] == BlankChoice {
		slog.Debug("The game has blank choices", "game", game.Debug())
		return ErrGameEmpty
	}
	// Choices in order, any blank at the end
	// No duplicate criteria
	lastChoice := game[0]
	for i := 1; i < MaxNumberOfChoicesPerGame; i++ {
		if lastChoice == BlankChoice {
			if game[i] != BlankChoice {
				return ErrGameChoiceOrderBlank
			}
			continue
		}
		if game[i] == BlankChoice {
			lastChoice = BlankChoice
			continue
		}
		if lastChoice >= game[i] {
			return ErrGameChoiceOrderCriterias
		}
		if game[i].Criteria().Id == lastChoice.Criteria().Id {
			return ErrGameRepatingCriteria
		}
		lastChoice = game[i]
	}
	// Unique solution
	if _, ok := game.Solve(); !ok {
		return ErrGameNoUniqueSolution
	}
	// No redundant card
	if (State{Game: game, mask: game.GetMask()}.HasRedundant()) {
		return ErrGameHasRedundant
	}
	return nil
}

// Returns the difficulty of this game.
// The game must be valid.
func (game Game) Difficulty() Difficulty {
	return max(game[0].Difficulty(), game[1].Difficulty(), game[2].Difficulty(),
		game[3].Difficulty(), game[4].Difficulty(), game[5].Difficulty())
}

// Returns true if the game has a unique solution
func (game Game) HasUniqueSolution() bool {
	return game.GetMask().Available() == 1
}

// Returns the solution to this game, if there is not one unique solution false
// will be returned.
// The game must be valid.
func (game Game) Solve() (code Code, ok bool) {
	mask := game.GetMask()
	return mask.GetCode(), mask.Available() == 1
}

// Returns the mask for this game
func (game Game) GetMask() CodeMask {
	return game[0].Mask().And(game[1].Mask()).And(game[2].Mask()).And(game[3].Mask()).And(game[4].Mask()).And(game[5].Mask())
}

// Returns a slice of criteria ids, a slice of verification cards
// with a random symbol, and a slice of laws (ids) for this game
func (game Game) GetCards() (criterias []int, verificationCards []string, laws []int) {
	l := game.NumberOfChoices()
	criterias = make([]int, l)
	laws = make([]int, l)
	vc := make([]VerificationCard, l)
	for i := range l {
		criterias[i] = int(game[i].Criteria().Id)
		vc[i] = game[i].Law().VerificationCard
		laws[i] = int(game[i].Law().Id)
	}
	verificationCards = getRandomVerificationSymbol(vc)
	return
}
