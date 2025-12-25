package main

import "sort"

func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Ints(happiness)
	var res int64 = 0
	count := 0
	for i := len(happiness) - 1; i >= len(happiness)-k; i-- {
		if happiness[i]-count > 0 {
			res += int64(happiness[i] - count)
		}
		count++
	}
	return res
}
