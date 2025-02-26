package main

import "slices"

func threeSum(nums []int) [][]int {
	// Sort the slice in ASC first
	slices.Sort(nums)
	var res [][]int

	for i, n := range nums {
		// Ignore duplicate nums, since we have already checked any combination starting with it before
		if i != 0 && n == nums[i-1] {
			continue
		}
		//          [1,2,3,4,5,6]
		// fix here _^ ^______^ do two pointers 2Sum
		// Doing 2Sum with two pointers, starting from the next idx
		l, r := i+1, len(nums)-1
		// Left can never exceed Right
		for l < r {
			threeSums := n + nums[l] + nums[r]
			if threeSums > 0 {
				// > 0, we can easily decrease the sum by moving r
				// since slice is sorted, it is guaranteed to be smaller once moved
				r--
			} else if threeSums < 0 {
				// < 0, we can easily increase the sum by moving l
				// since slice is sorted, it is guaranteed to be greater once moved
				l++
			} else {
				// Found the combination that equals to 0
				res = append(res, []int{n, nums[l], nums[r]})
				// Moving to next 2Sum left
				l++
				// Same as above, we ignore the same elements that have been checked
				for nums[l] == nums[l-1] && l < r {
					l++
				}
			}
		}
	}
	return res
}
