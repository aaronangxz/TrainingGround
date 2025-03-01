package main

func applyOperations(nums []int) []int {
	n, idx := len(nums), 0
	for i := 0; i < n; i++ {
		if i < n-1 && nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
		}
		if nums[i] != 0 {
			nums[idx], nums[i] = nums[i], nums[idx]
			idx++
		}
	}
	return nums
}
