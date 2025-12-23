package main

import "sort"

func maxTwoEvents(events [][]int) int {
	// Sort events by end time in descending order
	sort.Slice(events, func(i, j int) bool {
		return events[i][1] > events[j][1]
	})

	n := len(events)
	maxVals := make([]int, n+1) // To store max values from the end to the beginning
	maxVals[n] = 0              // Initialize the value after the last event as 0

	// Populate maxVals array with maximum values from right to left
	for i := n - 1; i >= 0; i-- {
		maxVals[i] = max(maxVals[i+1], events[i][2])
	}

	best := 0

	// Iterate over each event to find the best combination
	for i := 0; i < n; i++ {
		currentValue := events[i][2]
		// Binary search for the first event that ends before the current event starts
		idx := sort.Search(n, func(j int) bool {
			return events[j][1] < events[i][0]
		})
		if idx < n {
			currentValue += maxVals[idx]
		}
		best = max(best, currentValue)
	}

	return best
}
