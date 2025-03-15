package main

import "math"

func canStealKHouses(nums []int, k, capability int) bool {
	count := 0
	i := 0
	for i < len(nums) {
		if nums[i] <= capability {
			count++
			i += 2
		} else {
			i++
		}
	}
	return count >= k
}

func minCapability(nums []int, k int) int {
	left, right := 1, math.MaxInt32

	for left < right {
		mid := left + (right-left)/2
		if canStealKHouses(nums, k, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
