package main

func maximumDifference(nums []int) int {
	minVal := nums[0]
	maxDiff := -1
	for _, num := range nums[1:] {
		if num > minVal {
			diff := num - minVal
			if diff > maxDiff {
				maxDiff = diff
			}
		} else {
			minVal = num
		}
	}
	return maxDiff
}
