package main

func minOperations(nums []int) int {
	cnt := 0
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] == 0 {
			cnt++
			nums[i+1] = 1 - nums[i+1]
			nums[i+2] = 1 - nums[i+2]
		}
	}
	if nums[len(nums)-1] == 1 && nums[len(nums)-2] == 1 {
		return cnt
	}
	return -1
}
