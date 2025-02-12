package main

func maximumSum(nums []int) int {
	// Store the max sum of each index
	sumMap := make(map[int]int)

	// Fall back returns -1
	maxSum := -1
	for i := range nums {
		// Get digit sums
		digits := sumDigits(nums[i])

		// If seen before, compare the sum with the current index sum
		if v, ok := sumMap[digits]; ok {
			maxSum = max(maxSum, v+nums[i])
		}

		// Update the max sum
		sumMap[digits] = max(sumMap[digits], nums[i])
	}
	return maxSum
}

func sumDigits(n int) int {
	sum := 0

	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}
