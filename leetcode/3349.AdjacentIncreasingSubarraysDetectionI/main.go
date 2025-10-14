package main

func hasIncreasingSubarrays(nums []int, k int) bool {
	lastIncreasingSubarrayEnding := -1 // Ending index of last increasing subarray that is len >= k

	for left, right := 0, 1; right <= len(nums); right++ {
		if right == len(nums) || nums[right] <= nums[right-1] {
			if right-left >= 2*k { // Contains enough for at least two subsequent increasing subarrays
				return true
			} else if right-left >= k { // Is enough for at least one subarray
				if left == lastIncreasingSubarrayEnding { // lastIncreasingSubarrayEnding is adjacent
					return true
				}
				lastIncreasingSubarrayEnding = right
			}
			left = right
		}
	}

	return false
}
