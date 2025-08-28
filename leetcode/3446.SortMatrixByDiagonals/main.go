package main

import "sort"

func sortMatrix(grid [][]int) [][]int {
	n := len(grid)
	for i := 0; i < n; i++ {
		tmp := []int{}
		for j := 0; i+j < n; j++ {
			tmp = append(tmp, grid[i+j][j])
		}
		sort.Sort(sort.Reverse(sort.IntSlice(tmp)))
		for j := 0; i+j < n; j++ {
			grid[i+j][j] = tmp[j]
		}
	}

	for j := 1; j < n; j++ {
		tmp := []int{}
		for i := 0; j+i < n; i++ {
			tmp = append(tmp, grid[i][j+i])
		}
		sort.Ints(tmp)
		for i := 0; j+i < n; i++ {
			grid[i][j+i] = tmp[i]
		}
	}
	return grid
}
