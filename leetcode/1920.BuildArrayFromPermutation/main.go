package main

func buildArray(nums []int) []int {
	n := len(nums)

	for i := 0; i < n; i++ {
		nums[i] = nums[i] + (nums[nums[i]]%n)*n
	}

	for i := 0; i < n; i++ {
		nums[i] = nums[i] / n
	}

	return nums
}
