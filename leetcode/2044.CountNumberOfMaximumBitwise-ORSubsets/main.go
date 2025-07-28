package main

func countMaxOrSubsets(nums []int) int {
	maxOr := 0
	for _, n := range nums {
		maxOr |= n
	}

	var helper func(i, currOr int) int
	helper = func(i, currOr int) int {
		if i >= len(nums) {
			if currOr == maxOr {
				return 1
			}
			return 0
		}
		include := helper(i+1, currOr|nums[i])
		exclude := helper(i+1, currOr)
		return include + exclude
	}

	return helper(0, 0)
}
