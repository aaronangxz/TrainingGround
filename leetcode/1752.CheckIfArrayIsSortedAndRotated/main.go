package main

func check(nums []int) bool {
	inversions := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[(i+1)%len(nums)] {
			inversions++
		}
	}

	return inversions <= 1
}
