package main

import "fmt"

func findDiagonalOrder(mat [][]int) []int {
	if len(mat) == 0 {
		return nil
	}

	rowSize := len(mat)
	colSize := len(mat[0])

	i, j := 0, 0
	isMoveUpRight := true
	var res []int
	res = append(res, mat[i][j])

	for i < rowSize-1 || j < colSize-1 {
		if isMoveUpRight {
			//check boundary first
			if j == colSize-1 {
				//when it reaches the right most, move to the next row
				//and travel down-left next
				i++
				isMoveUpRight = !isMoveUpRight
			} else if i == 0 {
				//when it is on the top row
				//continue to move rightwards
				//and travel down-left next
				j++
				isMoveUpRight = !isMoveUpRight
			} else {
				//otherwise, just need to move up-right
				i--
				j++
			}
		} else {
			//check boundary first
			if i == rowSize-1 {
				//when it reaches the bottom row, move to next column
				//and travel up-right next
				j++
				isMoveUpRight = !isMoveUpRight
			} else if j == 0 {
				//when it is on the left most
				//continue to move downwards
				//and travel up-right next
				i++
				isMoveUpRight = !isMoveUpRight
			} else {
				//otherwise, just need to move down-left
				i++
				j--
			}
		}
		res = append(res, mat[i][j])
	}
	return res
}

func main() {
	mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(findDiagonalOrder(mat))
}
