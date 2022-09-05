package main

import (
	"github.com/aaronangxz/TrainingGround/common"
)

func findBuildings(heights []int) []int {
	maxHeight := 0
	var idx []int

	for i := len(heights) - 1; i >= 0; i-- {
		if heights[i] > maxHeight {
			idx = append(idx, i)
		}
		maxHeight = common.Max(maxHeight, heights[i])
	}
	common.ReverseSlice(&idx)
	return idx
}
