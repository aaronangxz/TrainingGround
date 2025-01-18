package main

import (
	"github.com/aaronangxz/TrainingGround/common"
)

func maximalSquare(matrix [][]byte) int {
	// Construct 2D matrix, fill with -1
	sqMap := make([][]int, len(matrix))
	for i := range sqMap {
		sqMap[i] = make([]int, len(matrix[0]))
	}
	for i := range sqMap {
		for j := range sqMap[i] {
			sqMap[i][j] = -1
		}
	}

	var dp func(r int, c int) int
	dp = func(r, c int) int {
		// Base case: out of bounds
		if r >= len(matrix) || c >= len(matrix[0]) {
			return 0
		}

		// If not cached
		if sqMap[r][c] == -1 {
			// Take the current cell as top left anchor,
			// check its bottom, right and diagonal
			bottom := dp(r+1, c)
			right := dp(r, c+1)
			diag := dp(r+1, c+1)

			// Only if current cell is '1',
			// update the count with 1 (self) + min of rest of square cells
			// Take min because if one is 0, then take 0 since it is not a square
			if string(matrix[r][c]) == "1" {
				sqMap[r][c] = 1 + common.Min(bottom, common.Min(right, diag))
			} else {
				// ignore '0' cells
				sqMap[r][c] = 0
			}
		}
		return sqMap[r][c]
	}

	dp(0, 0)
	maxSq := 0
	for i := range sqMap {
		for j := range sqMap[i] {
			// Find the max based on top left anchor cell so far
			maxSq = common.Max(maxSq, sqMap[i][j])
		}
	}
	// Return area
	return maxSq * maxSq
}
