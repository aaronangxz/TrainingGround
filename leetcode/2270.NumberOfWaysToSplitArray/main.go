package main

func waysToSplitArray(nums []int) int {
	// Keep track of total sum
	totalSum := 0
	for _, n := range nums {
		totalSum += n
	}

	left, res := 0, 0
	
	// Iterate until second last element
	// since it is not possible to split there
	for i := 0; i < len(nums)-1; i++ {
		// Continue to increment the left sum
		left += nums[i]
		// Right is simply total - left
		right := totalSum - left

		if left >= right {
			res++
		}
	}
	return res
}
