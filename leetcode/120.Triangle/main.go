package main

func minimumTotal(triangle [][]int) int {
	n := len(triangle)

	// Modify triangle in place from bottom up.
	for i := n - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			if triangle[i+1][j] < triangle[i+1][j+1] {
				triangle[i][j] += triangle[i+1][j]
			} else {
				triangle[i][j] += triangle[i+1][j+1]
			}
		}
	}

	// The top element now contains the minimum path sum.
	return triangle[0][0]
}
