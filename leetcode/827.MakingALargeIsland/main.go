package main

// Directions for DFS traversal (Right, Left, Down, Up)
var directions = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

// largestIsland finds the largest possible island by flipping one `0` to `1`
func largestIsland(grid [][]int) int {
	n := len(grid)
	islandID := 2 // Start labeling from 2
	islandSize := map[int]int{}

	// DFS function to explore and label an island
	var dfs func(int, int, int) int
	dfs = func(x, y, id int) int {
		// Boundary and visited check
		if x < 0 || y < 0 || x >= n || y >= n || grid[x][y] != 1 {
			return 0
		}
		grid[x][y] = id // Mark the cell with island ID
		size := 1
		for _, dir := range directions {
			size += dfs(x+dir[0], y+dir[1], id)
		}
		return size
	}

	// Step 1: Identify all islands and store their sizes
	maxIsland := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				islandSize[islandID] = dfs(i, j, islandID)
				maxIsland = max(maxIsland, islandSize[islandID])
				islandID++
			}
		}
	}

	// Step 2: Try flipping each `0` to `1` and calculate max possible island
	hasZero := false
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				hasZero = true
				seen := make(map[int]bool)
				newSize := 1 // Include the flipped cell
				for _, dir := range directions {
					ni, nj := i+dir[0], j+dir[1]
					if ni >= 0 && nj >= 0 && ni < n && nj < n {
						if id, exists := islandSize[grid[ni][nj]]; exists && !seen[grid[ni][nj]] {
							seen[grid[ni][nj]] = true
							newSize += id
						}
					}
				}
				maxIsland = max(maxIsland, newSize)
			}
		}
	}

	// If no `0` was found, return `n * n` (all 1s)
	if !hasZero {
		return n * n
	}
	return maxIsland
}
