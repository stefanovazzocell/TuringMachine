package store

import (
	"errors"
	"fmt"
	"log/slog"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/stefanovazzocell/TuringMachine/src/turingmachine/game"
)

const (
	// Maximum number of tries before GetRandomGameInRangeWithDifficulty gives
	// up.
	RandomGameMaxRetries = 100000
)

var (
	// Error returned when the Store file has the wrong size
	ErrInvalidFile = errors.New("the file size is not valid")
	// Error returned when we exceeded the threashold of retries for a random
	// game search
	ErrMaxRetries = errors.New("too many failed tries")
)

// A database for valid games
type Store struct {
	file *os.File
	// step indicates the count of all games with [0, (i+1)] choices.
	// ex: step[2] = how many games are there with 1 + 2 + 3 choices
	step [game.MaxNumberOfChoicesPerGame]int64
}

// Opens a game store
func OpenStore(filename string) (*Store, error) {
	// 1. Open the file
	store := &Store{}
	var err error
	store.file, err = os.Open(filename)
	if err != nil {
		return nil, err
	}
	// 2. Populate step
	start := time.Now()
	defer func() {
		slog.Info("store initialized", "duration", time.Since(start))
	}()
	return store, store.init()
}

// Creates (or overwrites) a game store
// Requires a 64-bit build
func CreateStore(filename string) (*Store, error) {
	err := solve(filename)
	if err != nil {
		return nil, err
	}
	return OpenStore(filename)
}

// Returns a debug string for the store
func (store *Store) Debug() string {
	sb := strings.Builder{}
	var start, end int64
	for i := range game.MaxNumberOfChoicesPerGame {
		start, end = store.GameRangeByChoices(i + 1)
		sb.WriteString(fmt.Sprintf("%d game(s) with %d criterias\n",
			end-start, i+1))
	}
	return sb.String()
}

// Returns true if the store contains a given game
func (store *Store) HasGame(g game.Game) (bool, error) {
	start, end := int64(0), store.NumberOfGames()

	var (
		sg  game.Game
		err error
	)

	for start < end {
		mid := (start + end) >> 1
		sg, err = store.GetGame(mid)
		if err != nil {
			return false, err
		}
		if sg == g {
			return true, nil
		}
		if sg.Value() < g.Value() {
			start = mid + 1
			continue
		}
		end = mid
	}
	sg, err = store.GetGame(start)
	return sg == g, err
}

// Returns the game at a given index
func (store *Store) GetGame(idx int64) (game.Game, error) {
	return game.GameFromReader(store.file, idx)
}

// Returns the range [start, end) of game indexes that have a given number of
// choices.
// Returns [0, 0] if an invalid number of choices was passed
func (store *Store) GameRangeByChoices(choices int) (start, end int64) {
	if choices == 0 || choices > game.MaxNumberOfChoicesPerGame {
		return
	}
	if choices > 1 {
		start = store.step[choices-2]
	}
	end = store.step[choices-1]
	return
}

// Returns the total number of games in this store
func (store *Store) NumberOfGames() int64 {
	return store.step[game.MaxNumberOfChoicesPerGame-1]
}

// Returns a random game in a given range [start, end)
func (store *Store) GetRandomGameInRange(start, end int64) (game.Game, error) {
	return store.GetGame(start + rand.Int63n(end-start))
}

// Returns a random game in a range with a given difficulty
func (store *Store) GetRandomGameInRangeWithDifficulty(start, end int64, difficulty game.Difficulty) (game.Game, error) {
	game, err := store.GetGame(start + rand.Int63n(end-start))
	maxTries := RandomGameMaxRetries
	for err == nil && game.Difficulty() != difficulty && maxTries > 0 {
		game, err = store.GetGame(start + rand.Int63n(end-start))
		maxTries--
	}
	if maxTries == 0 {
		err = ErrMaxRetries
	}
	return game, err
}

// Closes the store underlying file
func (store *Store) Close() error {
	return store.file.Close()
}

/*
* Helpers
**/

// Helper function to be called on create.
func (store *Store) init() error {
	nGames, err := store.numberOfGames()
	if err != nil {
		return err
	}
	store.step[5] = nGames
	for i := int64(4); i >= 0; i-- {
		store.step[i], err = store.firstGameInRangeWithChoices(int(i+2), 0, store.step[i+1])
		if err != nil {
			break
		}
		if store.step[i] == -1 {
			store.step[i] = 0
		}
	}
	return err
}

// Returns the index of the first game to have target choices within the range
// [start, end].
// If not found returns -1. On error getting a game, it returns -1 & the error.
func (store *Store) firstGameInRangeWithChoices(target int, start, end int64) (idx int64, err error) {
	var (
		mid     int64
		game    game.Game
		choices int
	)
	for start <= end {
		mid = (start + end) >> 1
		game, err = store.GetGame(mid)
		if err != nil {
			return -1, err
		}
		choices = game.NumberOfChoices()
		if choices == target {
			if mid == 0 {
				return mid, nil
			}
			// Check if the previous game is different
			game, err = store.GetGame(mid - 1)
			if err != nil {
				return -1, err
			}
			choices = game.NumberOfChoices()
			if choices != target {
				return mid, nil
			}
		} else if choices < target {
			start = mid + 1
			continue
		}
		end = mid - 1
	}
	return -1, nil
}

// Returns the total number of games
func (store *Store) numberOfGames() (int64, error) {
	info, err := store.file.Stat()
	if err != nil {
		return 0, err
	}
	size := info.Size()
	if size%game.MaxNumberOfChoicesPerGame != 0 {
		return 0, ErrInvalidFile
	}
	return size / game.MaxNumberOfChoicesPerGame, nil
}
