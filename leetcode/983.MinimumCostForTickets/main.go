package main

import (
	"github.com/aaronangxz/TrainingGround/common"
	"math"
)

type CostMap struct {
	Cost     int
	Duration int
}

func mincostTickets(days []int, costs []int) int {
	cMap := []CostMap{
		{Cost: costs[0], Duration: 1},
		{Cost: costs[1], Duration: 7},
		{Cost: costs[2], Duration: 30},
	}

	dp := make([]int, len(days)+1)

	// Start checking from the final day onwards
	for i := len(days) - 1; i >= 0; i-- {
		// Fixed ith day, this is the day the ticket is purchase
		// Moving jth day, this is to see the days where the ticket is valid till
		j := i
		dp[i] = math.MaxInt32

		// Attempt to test each ticket
		for _, c := range cMap {
			// Continue to move forward if the current day is still within
			// the ticket validity days
			for j < len(days) && days[j] < days[i]+c.Duration {
				j++
			}

			// Found the max day that the ticket can be stretched
			// Compare against the previously computed ticket cost on jth day
			// with the current cost combined
			dp[i] = common.Min(dp[i], c.Cost+dp[j])
		}
	}
	return dp[0]
}
