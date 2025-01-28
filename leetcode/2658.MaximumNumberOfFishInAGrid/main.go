package main

func findMaxFish(grid [][]int) int {
	R, C := len(grid), len(grid[0])
	maxFishes := 0

	// Adjacent neighbours' coordinates
	adj := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		// Already visited
		if grid[r][c] == 0 {
			return 0
		}

		// Get the current grid value first
		curr := grid[r][c]

		// Marking this grid as visited
		grid[r][c] = 0

		// DFS through the neighbours
		for _, cor := range adj {
			rNeigh, cNeigh := r+cor[0], c+cor[1]
			// Skip if: Out of bounds OR grid is land
			if rNeigh >= 0 && cNeigh >= 0 && rNeigh < R && cNeigh < C && grid[rNeigh][cNeigh] > 0 {
				// Increment the fish count of all the linked water cell
				curr += dfs(rNeigh, cNeigh)
			}
		}
		return curr
	}

	for r := range R {
		for c := range C {
			// Only DFS if it is not a land cell
			if grid[r][c] > 0 {
				maxFishes = max(maxFishes, dfs(r, c))
			}
		}
	}
	return maxFishes
}
