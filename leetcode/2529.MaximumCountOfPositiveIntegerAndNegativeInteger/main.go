package main

func maximumCount(nums []int) int {
	neg := binSearch(nums, 0)
	pos := len(nums) - binSearch(nums, 1)
	return max(neg, pos)
}

func binSearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
	result := len(nums)
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else {
			result = mid
			r = mid - 1
		}
	}
	return result
}
