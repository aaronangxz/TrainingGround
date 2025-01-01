package main

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	return binarySearch(nums, target, l, r)
}

func binarySearch(nums []int, target, l, r int) int {
	if l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		if target > nums[mid] {
			return binarySearch(nums, target, mid+1, r)
		}
		return binarySearch(nums, target, l, mid-1)
	}
	return -1
}
