package main

func countValidSelections(nums []int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}

	left := 0
	ans := 0

	for i := 0; i < len(nums); i++ {
		right := total - left

		if nums[i] == 0 {
			if left == right {
				ans += 2
			} else if abs(left-right) == 1 {
				ans += 1
			}
		}

		left += nums[i]
	}
	return ans
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
