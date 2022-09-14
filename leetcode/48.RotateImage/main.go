package main

import (
	"fmt"
	"github.com/aaronangxz/TrainingGround/common"
)

func main() {
	m := common.MakeSorted2DSlice(3, 3)
	rotate(m)
	fmt.Println(m)

	n := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(n)
	fmt.Println(n)
}

func rotate(matrix [][]int) {
	top, btm := 0, len(matrix)-1

	for top < btm {
		matrix[top], matrix[btm] = matrix[btm], matrix[top]
		top++
		btm--
	}

	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
