package main

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}

		if nums[l] <= nums[mid] {
			// we are in the left sorted part
			if target > nums[mid] || target < nums[l] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			// we are in the right sorted part
			if target < nums[mid] || target > nums[r] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}
