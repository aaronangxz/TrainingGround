package main

import (
	"fmt"
	"github.com/aaronangxz/TrainingGround/common"
	"math"
)

func maxSubArrayKadane(nums []int) int {
	localMax := 0
	globalMax := math.MinInt

	for _, n := range nums {
		//the largest sum so far
		//if current element is larger than the cumulative sum, means we are starting a new subarray
		localMax = common.Max(n, localMax+n)
		//the largest sum of contiguous array
		globalMax = common.Max(globalMax, localMax)
	}
	return globalMax
}

func main() {
	//s := common.MakeSlice(10)
	s := []int{-2, -1, -3, -4, -5, -6}
	fmt.Println(maxSubArrayKadane(s))
}
