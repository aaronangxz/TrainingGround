package MinimumOperationstoMakeColumnsStrictlyIncreasing

func minimumOperations(grid [][]int) int {
	ops := 0

	// matrix always has the same x,y, can use the first element's length as overall boundary
	for j := 0; j < len(grid[0]); j++ {
		for i := 1; i < len(grid); i++ {
			// If current cell is less than previous cell + 1
			if grid[i][j] < grid[i-1][j]+1 {
				// Total ops is the difference between the 2
				ops += grid[i-1][j] + 1 - grid[i][j]
				// Update the new value back for the next op
				grid[i][j] = grid[i-1][j] + 1
			}
		}
	}
	return ops
}
