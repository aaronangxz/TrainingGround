package main

import "sort"

// To return a merged interval with no overlapping intervals we can iterate through the intervals, comparing the current interval to the last merged interval.
// If the current intervals start is less or equal to the last merged intervals end then we update the merged interval with the max end of both intervals. Otherwise we add a new last merged interval.
// This requires the intervals to be sorted first making it a O(N * logN) time algorithm.
func merge(intervals [][]int) [][]int {
	// sort intervals ASC
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	// init merged intervals with first interval
	merged := [][]int{intervals[0]}

	for _, current := range intervals {
		// last merged interval is the latest interval in merged
		lastMerged := merged[len(merged)-1]

		// current start is within the last merged interval, update the end of last merged with max end
		if current[0] <= lastMerged[1] {
			lastMerged[1] = max(current[1], lastMerged[1])
		} else {
			// add new interval to the end of merged
			merged = append(merged, current)
		}
	}

	return merged
}
