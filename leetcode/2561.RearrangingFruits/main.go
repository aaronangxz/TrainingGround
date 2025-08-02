package main

import (
	"math"
	"sort"
)

func minCost(A, B []int) int64 {
	freq := map[int]int{}
	minVal := math.MaxInt
	for _, x := range A {
		freq[x]++
		if x < minVal {
			minVal = x
		}
	}
	for _, x := range B {
		freq[x]--
		if x < minVal {
			minVal = x
		}
	}

	extra := []int{}
	for k, v := range freq {
		if v%2 != 0 {
			return -1
		}
		for i := 0; i < abs(v)/2; i++ {
			extra = append(extra, k)
		}
	}

	sort.Ints(extra)
	var cost int64
	for i := 0; i < len(extra)/2; i++ {
		cost += int64(min(extra[i], 2*minVal))
	}
	return cost
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
