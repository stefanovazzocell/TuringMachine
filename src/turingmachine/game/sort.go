package game

// Sort the first 6 elements in a Game.
// Uses an optimal sorting network:
// https://en.wikipedia.org/wiki/Sorting_network#Optimal_sorting_networks
func SortGame6(game *Game) {
	// [(0,5),(1,3),(2,4)]
	if game[0] > game[5] {
		game[0], game[5] = game[5], game[0]
	}
	if game[1] > game[3] {
		game[1], game[3] = game[3], game[1]
	}
	if game[2] > game[4] {
		game[2], game[4] = game[4], game[2]
	}
	// [(1,2),(3,4)]
	if game[1] > game[2] {
		game[1], game[2] = game[2], game[1]
	}
	if game[3] > game[4] {
		game[3], game[4] = game[4], game[3]
	}
	// [(0,3),(2,5)]
	if game[0] > game[3] {
		game[0], game[3] = game[3], game[0]
	}
	if game[2] > game[5] {
		game[2], game[5] = game[5], game[2]
	}
	// [(0,1),(2,3),(4,5)]
	if game[0] > game[1] {
		game[0], game[1] = game[1], game[0]
	}
	if game[2] > game[3] {
		game[2], game[3] = game[3], game[2]
	}
	if game[4] > game[5] {
		game[4], game[5] = game[5], game[4]
	}
	// [(1,2),(3,4)]
	if game[1] > game[2] {
		game[1], game[2] = game[2], game[1]
	}
	if game[3] > game[4] {
		game[3], game[4] = game[4], game[3]
	}
}

// Sort the first 5 elements in a Game.
// Uses an optimal sorting network:
// https://en.wikipedia.org/wiki/Sorting_network#Optimal_sorting_networks
func SortGame5(game *Game) {
	// [(0,3),(1,4)]
	if game[0] > game[3] {
		game[0], game[3] = game[3], game[0]
	}
	if game[1] > game[4] {
		game[1], game[4] = game[4], game[1]
	}
	// [(0,2),(1,3)]
	if game[0] > game[2] {
		game[0], game[2] = game[2], game[0]
	}
	if game[1] > game[3] {
		game[1], game[3] = game[3], game[1]
	}
	// [(0,1),(2,4)]
	if game[0] > game[1] {
		game[0], game[1] = game[1], game[0]
	}
	if game[2] > game[4] {
		game[2], game[4] = game[4], game[2]
	}
	// [(1,2),(3,4)]
	if game[1] > game[2] {
		game[1], game[2] = game[2], game[1]
	}
	if game[3] > game[4] {
		game[3], game[4] = game[4], game[3]
	}
	// [(2,3)]
	if game[2] > game[3] {
		game[2], game[3] = game[3], game[2]
	}
}

// Sort the first 4 elements in a Game.
// Uses an optimal sorting network:
// https://en.wikipedia.org/wiki/Sorting_network#Optimal_sorting_networks
func SortGame4(game *Game) {
	// [(0,2),(1,3)]
	if game[0] > game[2] {
		game[0], game[2] = game[2], game[0]
	}
	if game[1] > game[3] {
		game[1], game[3] = game[3], game[1]
	}
	// [(0,1),(2,3)]
	if game[0] > game[1] {
		game[0], game[1] = game[1], game[0]
	}
	if game[2] > game[3] {
		game[2], game[3] = game[3], game[2]
	}
	// [(1,2)]
	if game[1] > game[2] {
		game[1], game[2] = game[2], game[1]
	}
}

// Sort the first 3 elements in a Game.
// Uses an optimal sorting network:
// https://en.wikipedia.org/wiki/Sorting_network#Optimal_sorting_networks
func SortGame3(game *Game) {
	// [(0,2)]
	if game[0] > game[2] {
		game[0], game[2] = game[2], game[0]
	}
	// [(0,1)]
	if game[0] > game[1] {
		game[0], game[1] = game[1], game[0]
	}
	// [(1,2)]
	if game[1] > game[2] {
		game[1], game[2] = game[2], game[1]
	}
}

// Sort the first 2 elements in a Game.
// Uses an optimal sorting network:
// https://en.wikipedia.org/wiki/Sorting_network#Optimal_sorting_networks
func SortGame2(game *Game) {
	// [(0,1)] -- trivial
	if game[0] > game[1] {
		game[0], game[1] = game[1], game[0]
	}
}
