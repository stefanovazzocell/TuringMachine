package store

// Merges multiple slice of game solutions (assuming those game slices are sorted)
func mergeSorted(solutions []solution) solution {
	// Size up the resulting slice to avoid multiple allocations
	size := 0
	for i := range solutions {
		size += len(solutions[i])
	}
	result := make(solution, size)
	// Initialize the result slice
	l := len(solutions) - 1
	lr := len(solutions[l]) // Length of result
	copy(result, solutions[l])
	// For each other solution, merge sorted
	for i := l - 1; i >= 0; i-- {
		// Indexes for the result and current solution slice
		ir := lr - 1
		is := len(solutions[i]) - 1
		// Index for the resulting slice
		idx := lr + is
		// Update expected lenght of the result slice
		lr += is + 1
		// Merge sort
		for ir >= 0 && is >= 0 {
			if result[ir].Value() < solutions[i][is].Value() {
				result[idx] = solutions[i][is]
				is--
			} else {
				result[idx] = result[ir]
				ir--
			}
			idx--
		}
		if is >= 0 {
			// Copy the remaining solutions
			copy(result, solutions[i][:is+1])
		}
	}
	return result
}
