package main

func countNegatives(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	r := 0
	c := cols - 1
	count := 0

	for r < rows && c >= 0 {
		if grid[r][c] < 0 {
			count += (rows - r)
			c--
		} else {
			r++
		}
	}
	return count
}
