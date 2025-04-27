package main

func countSubarrays(nums []int) int {
	result := 0
	for j := 1; j < len(nums)-1; j++ {
		if nums[j] == (nums[j-1]+nums[j+1])*2 {
			result++
		}
	}
	return result
}
