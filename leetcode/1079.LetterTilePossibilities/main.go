package main

func numTilePossibilities(tiles string) int {
	// Keeping track of the count of each unique tile
	countMap := make(map[string]int)
	for _, t := range tiles {
		countMap[string(t)]++
	}

	// Backtracking function
	var backtrack func() int
	backtrack = func() int {
		res := 0
		for k, v := range countMap {
			// Only if tile has remaining count
			if v > 0 {
				// Used one
				countMap[k]--
				// This is one possible combination
				res++
				// Continue to find the next combination (next level of tree)
				res += backtrack()
				// Revert for backtracking
				countMap[k]++
			}
		}
		return res
	}
	return backtrack()
}
