package game_test

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/stefanovazzocell/TuringMachine/src/turingmachine/game"
)

func TestGameReadWrite(t *testing.T) {
	t.Parallel()
	numGames := 1000

	for choices := 4; choices <= 6; choices++ {
		for range numGames {
			original, err := game.RandomSolvableGame(choices, game.HardDifficulty)
			if err != nil {
				t.Fatalf("Failed to generate random game: %v", err)
			}

			rawBuf := make([]byte, 6)
			original.WriteTo(rawBuf, 0)

			reader := bytes.NewReader(rawBuf)
			game, err := game.GameFromReader(reader, 0)
			if err != nil {
				t.Fatalf("[%s] written as %+d got error while reading: %v",
					original.Debug(), rawBuf, err)
			}
			if game != original {
				t.Fatalf("[%s] written as %+d was read as [%s]",
					original.Debug(), rawBuf, game.Debug())
			}
		}
	}
}

func TestGameGetCards(t *testing.T) {
	t.Parallel()
	numGames := 1000

	for choices := 4; choices <= 6; choices++ {
		for range numGames {
			original, err := game.RandomSolvableGame(choices, game.HardDifficulty)
			if err != nil {
				t.Fatalf("Failed to generate random game: %v", err)
			}
			t.Log(original.Debug())

			criterias, verificationCards, laws := original.GetCards()
			if len(criterias) != choices || len(verificationCards) != choices || len(laws) != choices {
				t.Fatalf("GetCards() returned (%+d, %+s, %+d) but expected only %d",
					criterias, verificationCards, laws, choices)
			}

			for i, criteria := range criterias {
				if criteria < 1 || game.NumberOfCriterias < criteria {
					t.Fatalf("Got invalid criterias %+d", criterias)
				}
				found := false
				for _, law := range game.Criterias[criteria-1].Laws {
					if int(law.Id) == laws[i] {
						found = true
						break
					}
				}
				if !found {
					t.Fatalf("GetCards() returned (%+d, %+s, %+d) but criteria %d has no matching law",
						criterias, verificationCards, laws, criteria)
				}
			}

			vc := make([]uint16, choices)
			for i := range choices {
				card := verificationCards[i]
				u64, err := strconv.ParseUint(card[len(card)-3:], 10, 16)
				if err != nil {
					t.Fatalf("GetCards() returned (%+d, %+s, %+d) but failed to parse verification card %s",
						criterias, verificationCards, laws, card)
				}
				vc[i] = uint16(u64)
			}

			crit := make([]uint8, choices)
			for i := range choices {
				crit[i] = uint8(criterias[i])
			}
			recovered, ok := game.GameFromCards(crit, vc)
			if !ok {
				t.Fatalf("GetCards() returned (%+d, %+s, %+d) - failed to convert it back to game",
					criterias, verificationCards, laws)
			}
			if recovered != original {
				t.Fatalf("GetCards() returned (%+d, %+s, %+d) - got mismatched recovered game [%s]",
					criterias, verificationCards, laws, recovered.Debug())
			}
		}
	}

}

func TestGameFromCards(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		criteriaCards     []uint8
		verificationCards []uint16
		isOk              bool
		validation        error
	}{
		{
			[]uint8{4, 9, 11, 14},
			[]uint16{447, 646, 566, 322},
			true,
			nil,
		},
		{
			[]uint8{4, 9, 11, 11},
			[]uint16{447, 646, 566, 322},
			false,
			nil,
		},
		{
			[]uint8{4, 9, 11, 11},
			[]uint16{447, 646, 461, 737},
			true,
			game.ErrGameRepatingCriteria,
		},
		{
			[]uint8{4, 9, 11, 11},
			[]uint16{447, 646, 566, 566},
			true,
			game.ErrGameChoiceOrderCriterias,
		},
		{
			[]uint8{14, 4, 9, 11},
			[]uint16{322, 447, 646, 566},
			true,
			nil,
		},
		{
			[]uint8{11, 22, 30, 33, 34, 40},
			[]uint16{287, 533, 389, 486, 547, 615},
			true,
			nil,
		},
		{
			[]uint8{4},
			[]uint16{447},
			true,
			game.ErrGameNoUniqueSolution,
		},
		{
			[]uint8{4, 9, 11, 14, 40},
			[]uint16{447, 646, 566, 322, 615},
			true,
			game.ErrGameHasRedundant,
		},
	}

	for i, testCase := range testCases {
		game, ok := game.GameFromCards(testCase.criteriaCards, testCase.verificationCards)
		if ok != testCase.isOk {
			t.Errorf("[%d] GameFromCards(%+d, %+d) = (%s, %v), want ok=%v",
				i, testCase.criteriaCards, testCase.verificationCards, game.Debug(), ok, testCase.isOk)
		}
		if !ok {
			continue
		}
		err := game.ValidateStrict()
		if !errors.Is(err, testCase.validation) {
			t.Errorf("[%d] (%s).ValidateStrict() = %v, but expected %v",
				i, game.Debug(), err, testCase.validation)
		}
	}
}

func TestGameSort(t *testing.T) {
	testCases := map[game.Game]game.Game{
		{}:                 {},
		{1}:                {1},
		{0, 0, 0, 0, 0, 1}: {1},
		{1, 2, 3, 4, 5, 6}: {1, 2, 3, 4, 5, 6},
		{0, 1, 2, 3, 4, 5}: {1, 2, 3, 4, 5, 0},
		{1, 2, 0, 3, 4, 5}: {1, 2, 3, 4, 5, 0},
		{1, 4, 2, 0, 3, 5}: {1, 2, 3, 4, 5, 0},
		{4, 1, 2, 0, 3, 5}: {1, 2, 3, 4, 5, 0},
		{4, 1, 2, 0, 3, 0}: {1, 2, 3, 4, 0, 0},
		{1, 2, 3, 4, 5, 0}: {1, 2, 3, 4, 5, 0},
	}

	for testCase, expected := range testCases {
		var original game.Game
		copy(original[:], testCase[:])
		testCase.Sort()
		if testCase != expected {
			t.Errorf("[%+d].sort() expected %+d, instead got %+d", original, expected, testCase)
		}
	}
}

func TestRandomSolvableGame(t *testing.T) {
	t.Parallel()
	numberOfGamesPerCombination := 10000
	difficulties := []game.Difficulty{game.HardDifficulty, game.StandardDifficulty, game.EasyDifficulty}

	for choices := 4; choices <= 6; choices++ {
		for _, difficulty := range difficulties {
			t.Run(fmt.Sprintf("%dchoices_%ddifficulty", choices, difficulty), func(t *testing.T) {
				t.Parallel()

				// Generate random games
				games := make([]game.Game, numberOfGamesPerCombination)
				var err error
				for i := range numberOfGamesPerCombination {
					games[i], err = game.RandomSolvableGame(choices, difficulty)
					if err != nil {
						t.Fatalf("Got error during generation: %v", err)
					}
				}

				// Strictly validate the games
				for i := range numberOfGamesPerCombination {
					err = games[i].ValidateStrict()
					if err != nil {
						t.Errorf("Game %s has failed validation: %v",
							games[i].Debug(), err)
					}
				}
			})
		}
	}
}

func TestGameString(t *testing.T) {
	t.Parallel()

	// Generate a large set of games.
	// It should always be the same set between runs of this test.
	i := (100000 * (game.MaxChoice + 1)) - 1
	testCases := make([]game.Game, i+1)
	for a := range byte(game.MaxChoice + 1) {
		for b := range byte(10) {
			for c := range byte(10) {
				for d := range byte(10) {
					for e := range byte(10) {
						for f := range byte(10) {
							testCases[i] = game.Game{
								game.Choice(a) % (game.MaxChoice + 1),
								game.Choice(b*2) % (game.MaxChoice + 1),
								game.Choice(c*3) % (game.MaxChoice + 1),
								game.Choice(d*5) % (game.MaxChoice + 1),
								game.Choice(e*7) % (game.MaxChoice + 1),
								game.Choice(f*11) % (game.MaxChoice + 1),
							}
							if !testCases[i].IsValid() {
								t.Fatalf("We should only generate valid games but got %s",
									testCases[i].Debug())
							}
							i--
						}
					}
				}
			}
		}
	}

	for _, expected := range testCases {
		// Try to convert the game into a string and back.
		// We should get the same game.
		asString := expected.String()
		actual, err := game.GameFromString(asString)
		if err != nil {
			t.Errorf("Got error decoding game string:\ngame: %s\nstring: %q\nerror: %v",
				expected.Debug(), asString, err)
		}
		if actual != expected {
			t.Errorf("Decoded game doesn't match original:\nexpected: %s\nstring: %q\nactual: %v",
				expected.Debug(), asString, actual.Debug())
		}
	}
}

func FuzzGameString(f *testing.F) {
	for a := range byte(game.MaxChoice + 1) {
		for b := range byte(2) {
			for c := range byte(2) {
				for d := range byte(2) {
					for e := range byte(2) {
						// skipping f
						for g := range byte(2) {
							f.Add(game.Game{
								game.Choice(a) % (game.MaxChoice + 1),
								game.Choice(b*2) % (game.MaxChoice + 1),
								game.Choice(c*3) % (game.MaxChoice + 1),
								game.Choice(d*5) % (game.MaxChoice + 1),
								game.Choice(e*7) % (game.MaxChoice + 1),
								game.Choice(g*11) % (game.MaxChoice + 1),
							}.String())
						}
					}
				}
			}
		}
	}

	f.Fuzz(func(t *testing.T, gameStr string) {
		firstPass, err := game.GameFromString(gameStr)
		if err != nil {
			t.SkipNow()
		}

		firstString := firstPass.String()
		secondPass, err := game.GameFromString(gameStr)
		if err != nil {
			t.Errorf("Got error decoding game string:\ngame: %s\nstring: %q\nerror: %v\ntest string: %q",
				firstPass.Debug(), firstString, err, gameStr)
		}
		if firstPass != secondPass {
			t.Errorf("Decoded game doesn't match original:\nfirstPass: %s\nstring: %q\nsecondPass: %v\ntest string: %q",
				firstPass.Debug(), firstString, secondPass.Debug(), gameStr)
		}
	})
}
