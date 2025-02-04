package main

func maxAscendingSum(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}

	// Keeping track of overall max and current subarray max
	maxSum := 0
	currMax := nums[0]
	for i := 1; i < len(nums); i++ {
		// Subarray breaks when it is not increasing
		if nums[i] <= nums[i-1] {
			// Update the overall max
			maxSum = max(maxSum, currMax)
			// Resets the current subarray max back to 0
			currMax = 0
		}
		// Increment current subarray max regardless
		currMax += nums[i]
	}
	// Account for the final subarray
	return max(maxSum, currMax)
}
