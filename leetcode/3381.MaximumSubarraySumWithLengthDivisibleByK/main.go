package main

import "math"

func maxSubarraySum(nums []int, k int) int64 {

	n := len(nums)
	subArraySumSizeK := make([]int64, n-k+1)
	curWindowSum := int64(0)
	for idx := 0; idx < k; idx++ {
		curWindowSum += int64(nums[idx])
	}
	subArraySumSizeK[0] = curWindowSum

	for idx := k; idx < n; idx++ {
		curWindowSum += int64(nums[idx] - nums[idx-k])
		subArraySumSizeK[idx-k+1] = curWindowSum
	}

	maxSum := int64(math.MinInt64)
	for idx := 0; idx < k; idx++ {
		currSum := int64(0)
		for idx2 := idx; idx2+k <= n; idx2 = idx2 + k {
			currSum = max(subArraySumSizeK[idx2], currSum+subArraySumSizeK[idx2])
			maxSum = max(maxSum, currSum)
		}
	}

	return maxSum

}
