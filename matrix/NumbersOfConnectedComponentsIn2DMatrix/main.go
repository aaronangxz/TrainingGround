package main

import "fmt"

func isSafe(matrix []string, row, col int, c uint8, n, l int, visited *[][]bool) bool {
	return row >= 0 && row < n && col >= 0 && col < l && matrix[row][col] == c && !(*visited)[row][col]
}

func dfs(matrix []string, row, col int, c uint8, n, l int, visited *[][]bool) {
	//4 directions to check from the current point
	//		   1, 0
	//	0,-1	X	  1, 0
	//		  -1, 0
	rowNumber := []int{-1, 1, 0, 0}
	colNumber := []int{0, 0, 1, -1}

	//mark as visited
	(*visited)[row][col] = true

	//4 directions to check
	for i := 0; i < 4; i++ {
		//only if the point we are going is:
		//1. Row not out of bounds
		//2. Col not out of bounds
		//3. Has the same element
		//4. Is not visited
		if isSafe(matrix, row+rowNumber[i], col+colNumber[i], c, n, l, visited) {
			dfs(matrix, row+rowNumber[i], col+colNumber[i], c, n, l, visited)
		}
	}
}

func connectedComponents(matrix []string) int {
	connected := 0
	n := len(matrix)
	l := len(matrix[0])

	//bool array to mark visited
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, l)
	}

	//travel across the whole matrix
	for i := 0; i < n; i++ {
		for j := 0; j < l; j++ {
			//only check if it is not visited
			if !visited[i][j] {
				c := matrix[i][j]
				dfs(matrix, i, j, c, n, l, &visited)
				connected++
				//at this point all connected blocks with this element c are marked as visited
			}
		}
	}
	return connected
}

func main() {
	fmt.Println(connectedComponents([]string{"bbba", "baaa"}))
	fmt.Println(connectedComponents([]string{"aabba", "aabba", "aaaca"}))
}
