package main

// Complexity:
// Time O(N) and Space O(1) where N is the length of nums.
func countSubarrays(nums []int, k int64) int64 {
	// We use the sliding window technique to compute the number of valid subarrays
	// for each end index.
	windowSize := int64(0)
	windowSum := int64(0)
	result := int64(0)

	for end := range nums {
		windowSize++
		windowSum += int64(nums[end])

		for windowSum*windowSize >= k {
			windowSize--
			windowSum -= int64(nums[end-int(windowSize)])
		}
		result += windowSize
	}
	return result
}
