package main

import "github.com/aaronangxz/TrainingGround/common"

func maxProfit(prices []int) int {
	maxGain := 0

	// Assume the first price is the lowest
	minPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		currPrice := prices[i]

		// The max gain is the max of current max gain and current price minus the lowest price
		maxGain = common.Max(maxGain, currPrice-minPrice)
		// go back and fill in the min price, in case the current price is the lowest
		minPrice = common.Min(minPrice, currPrice)
	}
	return maxGain
}
