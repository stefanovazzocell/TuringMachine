package game_test

import (
	"math/rand"
	"testing"

	"stefanovazzoler.com/turingmachine/src/turingmachine/game"
)

// Returns true if this state contains a redundant entry.
// It ignores games with a single choice
func hasRedundantSlow(state game.State) bool {
	choice := state.Game.NumberOfChoices()
	if choice <= 1 {
		return false
	}
	mask := state.Game.GetMask()
	var testGame game.Game
	for c := range choice {
		copy(testGame[:], state.Game[:])
		testGame[c] = 0
		if testGame.GetMask().Equal(mask) {
			return true
		}
	}
	return false
}

// Returns a random game
func randomGame() game.Game {
	gameRaw := [game.MaxNumberOfChoicesPerGame]byte{}
	u64 := rand.Uint64()
	max := byte(game.MaxChoice + 1)
	gameRaw[0] = byte(u64) % max
	gameRaw[1] = byte(u64>>8) % max
	gameRaw[2] = byte(u64>>16) % max
	gameRaw[3] = byte(u64>>24) % max
	gameRaw[4] = byte(u64>>32) % max
	gameRaw[5] = byte(u64>>40) % max
	return game.Game{game.Choice(gameRaw[0]), game.Choice(gameRaw[1]),
		game.Choice(gameRaw[2]), game.Choice(gameRaw[3]),
		game.Choice(gameRaw[4]), game.Choice(gameRaw[5])}
}

func TestHasRedundant(t *testing.T) {
	t.Parallel()
	numberOfTests := 10000000

	for range numberOfTests {
		g := randomGame()
		state := game.StateFromGame(g)
		iR := state.HasRedundant()
		iRS := hasRedundantSlow(state)
		if iR != iRS {
			t.Fatalf("[%s].HasRedundant() = %v but expected %v", g.Debug(), iR, iRS)
		}
	}
}

func BenchmarkHasRedundant(b *testing.B) {
	g, err := game.GameFromString("6D32H59CZ")
	if err != nil {
		b.Fatalf("error setting up game: %v", err)
	}
	state := game.StateFromGame(g)

	b.Run("Standard", func(b *testing.B) {
		for range b.N {
			_ = state.HasRedundant()
		}
	})
}
