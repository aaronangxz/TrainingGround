package main

type idxLoc struct {
	X int
	Y int
}

func firstCompleteIndex(arr []int, mat [][]int) int {
	posMap := make(map[int]idxLoc)

	// Keeping track of the current row, col fill status of each coordinates
	rowCnt := make([]int, len(mat))
	colCnt := make([]int, len(mat[0]))

	// Map the coordinates of the matrix for fast retrieval later
	for i := range len(mat) {
		for j := range len(mat[0]) {
			posMap[mat[i][j]] = idxLoc{X: i, Y: j}
		}
	}

	for i, a := range arr {
		c := posMap[a]
		// i.e. (1,2)
		// row 1 now has one more cell filled
		// col 2 now has one more cell filled
		rowCnt[c.X]++
		colCnt[c.Y]++

		// For row, compare with col count
		// For col, compare with row count
		// Because for it to be filled, a row need to contain the number of cell of that row
		// which is the total number of col in the matrix. vice versa
		if rowCnt[c.X] == len(mat[0]) || colCnt[c.Y] == len(mat) {
			return i
		}
	}
	return -1
}
