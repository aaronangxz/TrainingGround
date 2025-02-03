package main

func longestMonotonicSubarray(nums []int) int {
	// If array only has 1 element, there is only one subarray possible
	if len(nums) < 2 {
		return len(nums)
	}

	// Keeping track of the combined max, and the max for inc/desc respectively
	longest := 0
	inc := 1
	dec := 1

	// Starting from 1st index, so we can compare the previous element
	for i := 1; i < len(nums); i++ {
		// Found a descending sub-array, means increasing sub-array has broken
		if nums[i] < nums[i-1] {
			dec++
			inc = 1
		}
		// Found an increasing sub-array, means decreasing sub-array has broken
		if nums[i] > nums[i-1] {
			inc++
			dec = 1
		}
		// Equal, both sub-arrays have broken
		if nums[i] == nums[i-1] {
			inc = 1
			dec = 1
		}
		longest = max(longest, max(inc, dec))
	}
	return longest
}
