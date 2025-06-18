package main

import "sort"

func divideArray(nums []int, k int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums); i += 3 {
		a, b, c := nums[i], nums[i+1], nums[i+2]
		if c-a > k {
			return [][]int{}
		}
		res = append(res, []int{a, b, c})
	}
	return res
}
