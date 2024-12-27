package main

func findTargetSumWays(nums []int, target int) int {
	// Calculate total of nums for array construction
	totalSum := 0
	for _, n := range nums {
		totalSum += n
	}

	// 2D dp slice - [index][total]
	dp := make([][]int, len(nums))
	// MIN placeholder to indicate no values
	MIN := 0 - totalSum - 1

	// Fill dp slice, need to initialize the nested slice inside
	for i := range dp {
		// Fill with min value
		// 2*totalSum to handle negative range [-totalSum,totalSum]
		dp[i] = make([]int, 2*totalSum+1)
		for j := 0; j < 2*totalSum+1; j++ {
			dp[i][j] = MIN
		}
	}

	var backtrack func(int, int) int
	backtrack = func(currentIdx, currentTotal int) int {
		// Base case: If we reach the last index, means there is 1 result
		if currentIdx == len(nums) {
			if currentTotal == target {
				return 1
			} else {
				return 0
			}
		}

		// Memoized before, return from dp slice
		if dp[currentIdx][currentTotal+totalSum] != MIN {
			return dp[currentIdx][currentTotal+totalSum]
		}

		// Else, calculate the left (Additon) and right (Subtraction) branch
		// The sum is the total possible ways thus far
		dp[currentIdx][currentTotal+totalSum] = backtrack(currentIdx+1, currentTotal+nums[currentIdx]) + backtrack(currentIdx+1, currentTotal-nums[currentIdx])
		return dp[currentIdx][currentTotal+totalSum]
	}

	// Start from 0 index, 0 sum (Root)
	return backtrack(0, 0)
}
