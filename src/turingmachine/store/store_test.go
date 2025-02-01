package store_test

import (
	"crypto/sha256"
	"fmt"
	"io"
	"math/rand"
	"os"
	"slices"
	"sync"
	"testing"

	"github.com/stefanovazzocell/TuringMachine/src/turingmachine/game"
	"github.com/stefanovazzocell/TuringMachine/src/turingmachine/store"
)

// Returns a random file name in the current directory
func tmpFile() string {
	return fmt.Sprintf("./test_%d.tmp", rand.Uint64())
}

// Attempts to cleanup a file
func rmFile(filename string) {
	_ = os.Remove(filename)
}

func hashFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

// Returns true if two files are equal
func hashEqual(filenameA, filenameB string) (bool, error) {
	var hashA, hashB []byte
	var errA, errB error
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		hashA, errA = hashFile(filenameA)
		wg.Done()
	}()
	hashB, errB = hashFile(filenameB)
	wg.Wait()

	if errA != nil {
		return false, errA
	}
	if errB != nil {
		return false, errB
	}
	return slices.Equal(hashA, hashB), nil
}

// Returns a store, a cleanup function, and an error
func getStore() (s *store.Store, cleanupFn func(), err error) {
	filename := tmpFile()
	s, err = store.CreateStore(filename)
	cleanupFn = func() {
		_ = s.Close()
		rmFile(filename)
	}
	return
}

func validateAllGamesInStore(s *store.Store, t *testing.T) {
	nGames := s.NumberOfGames()
	for gIdx := range nGames {
		g, err := s.GetGame(gIdx)
		if err != nil {
			t.Fatalf("Failed to read game %d: %v", gIdx, err)
		}
		err = g.ValidateStrict()
		if err != nil {
			t.Fatalf("Game %+d (%s) is not valid: %v", g, g.Debug(), err)
		}
	}
}

func spotCheckGamesInStoreByDifficulty(s *store.Store, t *testing.T, difficulty game.Difficulty, choices int, n int) {
	var found bool
	for range n {
		g, err := game.RandomSolvableGame(choices, difficulty)
		if err != nil {
			t.Fatalf("Error from RandomSolvableGame(%d, %d): %v",
				choices, difficulty, err)
		}
		found, err = s.HasGame(g)
		if err != nil {
			t.Fatalf("Errored when trying to find game %+d (%s) in store: %v",
				g, g.Debug(), err)
		}
		if !found {
			t.Fatalf("Game %+d (%s) not found in store", g, g.Debug())
		}
	}
}

func TestStore(t *testing.T) {
	expectedGames := int64(20881150)

	// Setup
	filename := tmpFile()
	defer rmFile(filename)
	// Create stores
	store, err := store.CreateStore(filename)
	if err != nil {
		t.Fatalf("Failed to create first store: %v", err)
	}

	nGames := store.NumberOfGames()
	if nGames != expectedGames {
		// At least 1 game expected
		t.Fatalf("Expected %d games, instead got %d",
			expectedGames, nGames)
	}

	t.Run("ValidateAllGames", func(t *testing.T) {
		t.Parallel()
		validateAllGamesInStore(store, t)
	})

	nChecks := 100
	for choices := 4; choices <= 6; choices++ {
		func(choices int) {
			t.Run(fmt.Sprintf("%dchoices", choices), func(t *testing.T) {
				t.Parallel()

				t.Run("SpotCheckEasy", func(t *testing.T) {
					t.Parallel()
					spotCheckGamesInStoreByDifficulty(store, t, game.EasyDifficulty, choices, nChecks)
				})
				t.Run("SpotCheckStandard", func(t *testing.T) {
					t.Parallel()
					spotCheckGamesInStoreByDifficulty(store, t, game.StandardDifficulty, choices, 2*nChecks)
				})
				t.Run("SpotCheckHard", func(t *testing.T) {
					t.Parallel()
					spotCheckGamesInStoreByDifficulty(store, t, game.HardDifficulty, choices, 3*nChecks)
				})
			})
		}(choices)
	}
}

func TestCreateStore(t *testing.T) {
	// Setup
	filenameA, filenameB := tmpFile(), tmpFile()
	defer rmFile(filenameA)
	defer rmFile(filenameB)
	// Create stores
	storeA, err := store.CreateStore(filenameA)
	if err != nil {
		t.Fatalf("Failed to create first store: %v", err)
	}
	storeB, err := store.CreateStore(filenameB)
	if err != nil {
		storeA.Close()
		t.Fatalf("Failed to create second store: %v", err)
	}
	// Try to retrieve some stats
	startB, endB := storeB.GameRangeByChoices(0)
	if startB+endB != 0 {
		t.Errorf("There should be more than 0 games with no choices, instead got %d",
			endB-startB)
	}
	startB, endB = storeB.GameRangeByChoices(game.MaxNumberOfChoicesPerGame + 1)
	if startB+endB != 0 {
		t.Errorf("There should be more than 0 games with %d choices, instead got %d",
			game.MaxNumberOfChoicesPerGame+1, endB-startB)
	}
	startB, endB = storeB.GameRangeByChoices(game.MaxNumberOfChoicesPerGame)
	if endB == startB {
		t.Errorf("There should be more than %d games with %d choices",
			endB-startB, game.MaxNumberOfChoicesPerGame)
	}
	// Close stores and compare
	if err = storeA.Close(); err != nil {
		t.Errorf("Failed to close storeA: %v", err)
	}
	if err = storeB.Close(); err != nil {
		t.Errorf("Failed to close storeB: %v", err)
	}
	ok, err := hashEqual(filenameA, filenameB)
	if err != nil {
		t.Fatalf("Failed to compare hashes: %v", err)
	}
	if !ok {
		t.Fatal("Hashes don't match")
	}
	// Try to open one of the two
	storeX, err := store.OpenStore(filenameA)
	if err != nil {
		t.Fatalf("Failed to open store: %v", err)
	}
	startX, endX := storeX.GameRangeByChoices(game.MaxNumberOfChoicesPerGame)
	if startX != startB || endX != endB {
		t.Errorf("Expected [%d, %d) games with %d choices, instead got [%d, %d)",
			startB, endB, game.MaxNumberOfChoicesPerGame, startX, endX)
	}
	g, err := storeX.GetRandomGameInRange(startX, endX)
	if err != nil || !g.IsValid() || g.NumberOfChoices() != game.MaxNumberOfChoicesPerGame {
		t.Errorf("Expected a valid game, instead got (%s, %v)",
			g.Debug(), err)
	}
	if err = storeX.Close(); err != nil {
		t.Errorf("Failed to close storeX: %v", err)
	}
}

func BenchmarkGet(b *testing.B) {
	store, cleanup, err := getStore()
	defer cleanup()
	if err != nil {
		b.Fatal(err)
	}
	start, games := store.GameRangeByChoices(game.MaxNumberOfChoicesPerGame)
	aGame, err := store.GetGame(games >> 1)
	if err != nil {
		b.Fatal(err)
	}

	b.Run("Sequential", func(b *testing.B) {
		b.Run("GetGame", func(b *testing.B) {
			var err error
			for i := range int64(b.N) {
				_, err = store.GetGame(i % games)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
		b.Run("GetRandomGameInRange", func(b *testing.B) {
			var err error
			for range b.N {
				_, err = store.GetRandomGameInRange(start, games)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
		b.Run("GetRandomGameInRangeWithDifficulty", func(b *testing.B) {
			b.Run("Hard", func(b *testing.B) {
				var err error
				for range b.N {
					_, err = store.GetRandomGameInRangeWithDifficulty(start, games, game.HardDifficulty)
					if err != nil {
						b.Fatal(err)
					}
				}
			})
			b.Run("Standard", func(b *testing.B) {
				var err error
				for range b.N {
					_, err = store.GetRandomGameInRangeWithDifficulty(start, games, game.StandardDifficulty)
					if err != nil {
						b.Fatal(err)
					}
				}
			})
			b.Run("Easy", func(b *testing.B) {
				var err error
				for range b.N {
					_, err = store.GetRandomGameInRangeWithDifficulty(start, games, game.EasyDifficulty)
					if err != nil {
						b.Fatal(err)
					}
				}
			})
		})
		// For comparison to non-store systems
		b.Run("RandomSolvableGame", func(b *testing.B) {
			b.Run("Hard", func(b *testing.B) {
				for range b.N {
					game.RandomSolvableGame(6, game.HardDifficulty)
				}
			})
			b.Run("Standard", func(b *testing.B) {
				for range b.N {
					game.RandomSolvableGame(6, game.StandardDifficulty)
				}
			})
			b.Run("Easy", func(b *testing.B) {
				for range b.N {
					game.RandomSolvableGame(6, game.EasyDifficulty)
				}
			})
		})
		b.Run("ValidateStrict", func(b *testing.B) {
			for range b.N {
				err = aGame.ValidateStrict()
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	})
	b.Run("Parallel", func(b *testing.B) {
		b.Run("GetGame", func(b *testing.B) {
			b.RunParallel(func(p *testing.PB) {
				for p.Next() {
					_, err = store.GetGame(games / 2)
					if err != nil {
						b.Fatal(err)
					}
				}
			})
		})
		b.Run("GetRandomGameInRange", func(b *testing.B) {
			b.RunParallel(func(p *testing.PB) {
				for p.Next() {
					_, err = store.GetRandomGameInRange(start, games)
					if err != nil {
						b.Fatal(err)
					}
				}
			})
		})
		b.Run("GetRandomGameInRangeWithDifficulty", func(b *testing.B) {
			b.Run("Hard", func(b *testing.B) {
				b.RunParallel(func(p *testing.PB) {
					for p.Next() {
						_, err = store.GetRandomGameInRangeWithDifficulty(start, games, game.HardDifficulty)
						if err != nil {
							b.Fatal(err)
						}
					}
				})
			})
			b.Run("Standard", func(b *testing.B) {
				b.RunParallel(func(p *testing.PB) {
					for p.Next() {
						_, err = store.GetRandomGameInRangeWithDifficulty(start, games, game.StandardDifficulty)
						if err != nil {
							b.Fatal(err)
						}
					}
				})
			})
			b.Run("Easy", func(b *testing.B) {
				b.RunParallel(func(p *testing.PB) {
					for p.Next() {
						_, err = store.GetRandomGameInRangeWithDifficulty(start, games, game.EasyDifficulty)
						if err != nil {
							b.Fatal(err)
						}
					}
				})
			})
		})
		// For comparison to non-store systems
		b.Run("RandomSolvableGame", func(b *testing.B) {
			b.Run("Hard", func(b *testing.B) {
				b.RunParallel(func(p *testing.PB) {
					for p.Next() {
						game.RandomSolvableGame(6, game.HardDifficulty)
					}
				})
			})
			b.Run("Standard", func(b *testing.B) {
				b.RunParallel(func(p *testing.PB) {
					for p.Next() {
						game.RandomSolvableGame(6, game.StandardDifficulty)
					}
				})
			})
			b.Run("Easy", func(b *testing.B) {
				b.RunParallel(func(p *testing.PB) {
					for p.Next() {
						game.RandomSolvableGame(6, game.EasyDifficulty)
					}
				})
			})
		})
		b.Run("ValidateStrict", func(b *testing.B) {
			b.RunParallel(func(p *testing.PB) {
				for p.Next() {
					err = aGame.ValidateStrict()
					if err != nil {
						b.Fatal(err)
					}
				}
			})
		})
	})
}
