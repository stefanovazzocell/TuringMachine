package game_test

import (
	"slices"
	"testing"

	"github.com/stefanovazzocell/TuringMachine/src/turingmachine/game"
)

func TestSortGame(t *testing.T) {
	t.Parallel()

	nGames := 10000000

	for range nGames {
		g := randomGame()

		for n := 1; n <= 6; n++ {
			slicesSort := g
			slices.Sort(slicesSort[:n])

			netSort := g
			switch n {
			case 6:
				game.SortGame6(&netSort)
			case 5:
				game.SortGame5(&netSort)
			case 4:
				game.SortGame4(&netSort)
			case 3:
				game.SortGame3(&netSort)
			case 2:
				game.SortGame2(&netSort)
			}

			if slicesSort != netSort {
				t.Fatalf("[%s] should be sorted as [%s], instead sorted as [%s]",
					g, slicesSort, netSort)
			}
		}
	}
}

func BenchmarkSortGame(b *testing.B) {
	nGames := 1000
	games := make([]game.Game, nGames)
	for i := range nGames {
		games[i] = randomGame()
	}

	b.Run("6", func(b *testing.B) {
		choices := 6
		b.Run("SlicesSort", func(b *testing.B) {
			var i int
			var g game.Game

			for i = range b.N {
				g = games[i%nGames]
				slices.Sort(g[:choices])
			}
		})
		b.Run("SortGame", func(b *testing.B) {
			var i int
			var g game.Game

			for i = range b.N {
				g = games[i%nGames]
				game.SortGame6(&g)
			}
		})
		b.Run("SortGameIf", func(b *testing.B) {
			var i int
			var g game.Game

			for i = range b.N {
				g = games[i%nGames]
				if choices == 6 {
					game.SortGame6(&g)
				} else if choices == 5 {
					game.SortGame5(&g)
				} else if choices == 4 {
					game.SortGame4(&g)
				} else if choices == 3 {
					game.SortGame3(&g)
				} else if choices == 2 {
					game.SortGame2(&g)
				}
			}
		})
	})

	b.Run("5", func(b *testing.B) {
		choices := 5
		b.Run("SlicesSort", func(b *testing.B) {
			var i int
			var g game.Game

			for i = range b.N {
				g = games[i%nGames]
				slices.Sort(g[:choices])
			}
		})
		b.Run("SortGame", func(b *testing.B) {
			var i int
			var g game.Game

			for i = range b.N {
				g = games[i%nGames]
				game.SortGame5(&g)
			}
		})
		b.Run("SortGameIf", func(b *testing.B) {
			var i int
			var g game.Game

			for i = range b.N {
				g = games[i%nGames]
				if choices == 6 {
					game.SortGame6(&g)
				} else if choices == 5 {
					game.SortGame5(&g)
				} else if choices == 4 {
					game.SortGame4(&g)
				} else if choices == 3 {
					game.SortGame3(&g)
				} else if choices == 2 {
					game.SortGame2(&g)
				}
			}
		})
	})

	b.Run("4", func(b *testing.B) {
		choices := 4
		b.Run("SlicesSort", func(b *testing.B) {
			var i int
			var g game.Game

			for i = range b.N {
				g = games[i%nGames]
				slices.Sort(g[:choices])
			}
		})
		b.Run("SortGame", func(b *testing.B) {
			var i int
			var g game.Game

			for i = range b.N {
				g = games[i%nGames]
				game.SortGame4(&g)
			}
		})
		b.Run("SortGameIf", func(b *testing.B) {
			var i int
			var g game.Game

			for i = range b.N {
				g = games[i%nGames]
				if choices == 6 {
					game.SortGame6(&g)
				} else if choices == 5 {
					game.SortGame5(&g)
				} else if choices == 4 {
					game.SortGame4(&g)
				} else if choices == 3 {
					game.SortGame3(&g)
				} else if choices == 2 {
					game.SortGame2(&g)
				}
			}
		})
	})

	b.Run("3", func(b *testing.B) {
		choices := 3
		b.Run("SlicesSort", func(b *testing.B) {
			var i int
			var g game.Game

			for i = range b.N {
				g = games[i%nGames]
				slices.Sort(g[:choices])
			}
		})
		b.Run("SortGame", func(b *testing.B) {
			var i int
			var g game.Game

			for i = range b.N {
				g = games[i%nGames]
				game.SortGame3(&g)
			}
		})
		b.Run("SortGameIf", func(b *testing.B) {
			var i int
			var g game.Game

			for i = range b.N {
				g = games[i%nGames]
				if choices == 6 {
					game.SortGame6(&g)
				} else if choices == 5 {
					game.SortGame5(&g)
				} else if choices == 4 {
					game.SortGame4(&g)
				} else if choices == 3 {
					game.SortGame3(&g)
				} else if choices == 2 {
					game.SortGame2(&g)
				}
			}
		})
	})

	b.Run("2", func(b *testing.B) {
		choices := 2
		b.Run("SlicesSort", func(b *testing.B) {
			var i int
			var g game.Game

			for i = range b.N {
				g = games[i%nGames]
				slices.Sort(g[:choices])
			}
		})
		b.Run("SortGame", func(b *testing.B) {
			var i int
			var g game.Game

			for i = range b.N {
				g = games[i%nGames]
				game.SortGame2(&g)
			}
		})
		b.Run("SortGameIf", func(b *testing.B) {
			var i int
			var g game.Game

			for i = range b.N {
				g = games[i%nGames]
				if choices == 6 {
					game.SortGame6(&g)
				} else if choices == 5 {
					game.SortGame5(&g)
				} else if choices == 4 {
					game.SortGame4(&g)
				} else if choices == 3 {
					game.SortGame3(&g)
				} else if choices == 2 {
					game.SortGame2(&g)
				}
			}
		})
	})

	b.Run("1", func(b *testing.B) {
		choices := 1
		b.Run("SlicesSort", func(b *testing.B) {
			var i int
			var g game.Game

			for i = range b.N {
				g = games[i%nGames]
				slices.Sort(g[:choices])
			}
		})
		b.Run("SortGame", func(b *testing.B) {
			var i int
			var g game.Game

			for i = range b.N {
				g = games[i%nGames]
				// N/A
				_ = g
			}
		})
		b.Run("SortGameIf", func(b *testing.B) {
			var i int
			var g game.Game

			for i = range b.N {
				g = games[i%nGames]
				if choices == 6 {
					game.SortGame6(&g)
				} else if choices == 5 {
					game.SortGame5(&g)
				} else if choices == 4 {
					game.SortGame4(&g)
				} else if choices == 3 {
					game.SortGame3(&g)
				} else if choices == 2 {
					game.SortGame2(&g)
				}
			}
		})
	})
}
