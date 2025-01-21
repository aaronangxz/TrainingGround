package main

import "math"

func gridGame(grid [][]int) int64 {
	n := len(grid[0])
	prefixSumTop, prefixSumBtm := make([]int, n), make([]int, n)

	// Populate prefix sum slices for row1 and row2
	prefixSumTop[0] = grid[0][0]
	prefixSumBtm[0] = grid[1][0]

	for i := 1; i < n; i++ {
		prefixSumTop[i] = prefixSumTop[i-1] + grid[0][i]
		prefixSumBtm[i] = prefixSumBtm[i-1] + grid[1][i]
	}

	// Set ans as maximum
	res := math.MaxInt64

	for i := range grid[0] {
		// Simulate top row
		// Assuming the path is from current idx to last idx
		top := prefixSumTop[n-1] - prefixSumTop[i]

		// Simulate bottom row
		// Assuming the path is from idx 0 to current idx
		// * -> Path of First bot (Collect points on the way)
		// x -> Path of Second bot with points (To collect the remaining points left)
		// Goal: Find the max points possible for second bot from top or bottom row

		// i = 0
		// * x x x x -> Find from 1 to 4 on top row
		// * * * * * -> None on bottom row
		// First bot goes from (0,0) -> (1,0) -> (1,1) -> (1,2) -> (1,3) -> (1,4)

		// i = 1
		// * * x x x -> Find from 2 to 4 on top row
		// x * * * * -> Find from 0 to 0 on bottom row
		// First bot goes from (0,0) -> (0,1) -> (1,1) -> (1,2) -> (1,3) -> (1,4)

		// i = 2
		// * * * x x -> Find from 3 to 4 on top row
		// x x * * * -> Find from 0 to 1 on bottom row
		// First bot goes from (0,0) -> (0,1) -> (0,2) -> (1,2) -> (1,3) -> (1,4)

		// i = 3
		// * * * * x -> Find from 4 to 4 on top row
		// x x x * * -> Find from 0 to 2 on bottom row
		// First bot goes from (0,0) -> (0,1) -> (0,2) -> (0,3) -> (1,3) -> (1,4)

		// i = 4
		// * * * * * -> None on top row
		// x x x x * -> Find from 0 to 3 on bottom row
		// First bot goes from (0,0) -> (0,1) -> (0,2) -> (0,3) -> (0,4) -> (1,4)

		var btm int
		if i > 0 {
			btm = prefixSumBtm[i-1]
		} else {
			btm = 0
		}

		// Second bot always wants to achieve the max
		secondBot := max(top, btm)
		// But we want to find the min that second bot can achieve
		res = min(res, secondBot)
	}
	return int64(res)
}
