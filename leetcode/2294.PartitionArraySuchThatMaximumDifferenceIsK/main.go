package main

import "sort"

func partitionArray(nums []int, k int) int {
	sort.Ints(nums)
	ans := 1
	l := 0

	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[l] > k {
			ans++
			l = i
		}
	}

	return ans
}
