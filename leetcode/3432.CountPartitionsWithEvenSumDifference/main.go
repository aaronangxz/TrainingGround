package main

func countPartitions(nums []int) int {
	var total int64 = 0
	for _, x := range nums {
		total += int64(x)
	}
	if total%2 != 0 {
		return 0
	}
	return len(nums) - 1
}
