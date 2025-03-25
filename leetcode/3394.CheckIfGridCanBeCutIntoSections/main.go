package main

import "sort"

func checkValidCuts(n int, rectangles [][]int) bool {
	var xIntervals, yIntervals [][]int

	for _, rect := range rectangles {
		xIntervals = append(xIntervals, []int{rect[0], rect[2]})
		yIntervals = append(yIntervals, []int{rect[1], rect[3]})
	}

	return check(xIntervals) || check(yIntervals)
}

func check(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	sections := 0
	maxEnd := intervals[0][1]

	for _, interval := range intervals {
		if maxEnd <= interval[0] {
			sections++
		}
		maxEnd = max(maxEnd, interval[1])
	}

	return sections >= 2
}
