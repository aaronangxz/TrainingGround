package main

func canPartition(nums []int) bool {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	// If total sum is odd, can't partition into equal subsets
	if totalSum%2 != 0 {
		return false
	}

	target := totalSum / 2
	dp := make([]bool, target+1)
	dp[0] = true // Base case: sum 0 can always be achieved

	for _, num := range nums {
		// Iterate backwards to avoid overwriting values we need to check
		for j := target; j >= num; j-- {
			if dp[j-num] {
				dp[j] = true
			}
		}
	}

	return dp[target]
}
