package main

import "sort"

func maxValue(events [][]int, k int) int {
	// nlogn, sort the events
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	// nlogn, find the first event that can be started after the current one
	nextValid := make([]int, len(events))
	for i := 0; i < len(events); i++ {
		nextValid[i] = sort.Search(len(events), func(j int) bool {
			return events[j][0] > events[i][1]
		})
	}

	// n * k, knapsack top down dp, each time we take and jump to the nexr valid even, or simply skip
	memo := make(map[[2]int]int)
	var dfs func(idx int, have int) int
	dfs = func(idx int, have int) int {
		if idx >= len(events) || have == k {
			return 0
		}
		if v, ok := memo[[2]int{idx, have}]; ok {
			return v
		}

		skip := dfs(idx+1, have)
		take := events[idx][2] + dfs(nextValid[idx], have+1)

		memo[[2]int{idx, have}] = max(skip, take)
		return memo[[2]int{idx, have}]
	}

	return dfs(0, 0)
}
