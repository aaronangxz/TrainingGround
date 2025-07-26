package main

import "slices"

func maxSubarrays(n int, conflictingPairs [][]int) int64 {
	// Ensure conflictingPairs[i][0] is always the smaller value
	for i := range conflictingPairs {
		if conflictingPairs[i][0] > conflictingPairs[i][1] {
			conflictingPairs[i][0], conflictingPairs[i][1] = conflictingPairs[i][1], conflictingPairs[i][0]
		}
	}

	// Sort conflictingPairs based on the first element of each pair
	slices.SortFunc(conflictingPairs, func(a, b []int) int { return b[0] - a[0] })

	rightBoundary := n + 1

	conflictIdx := 0
	largestCorrection := 0
	var validSubarrays int64
	var curPenalty, maxPenalty int64

	// Traverse from n down to 1
	for b := n; b > 0; b-- {
		// Process conflicting pairs affecting `b`
		for ; conflictIdx < len(conflictingPairs) && conflictingPairs[conflictIdx][0] >= b; conflictIdx++ {
			if rightBoundary > conflictingPairs[conflictIdx][1] {
				// Found a new correction point
				largestCorrection = rightBoundary - conflictingPairs[conflictIdx][1]
				rightBoundary = conflictingPairs[conflictIdx][1]
				curPenalty = 0
			} else {
				// Minimize the correction penalty
				largestCorrection = min(largestCorrection, conflictingPairs[conflictIdx][1]-rightBoundary)
			}
		}

		validSubarrays += int64(rightBoundary - b)
		curPenalty += int64(largestCorrection)
		maxPenalty = max(maxPenalty, curPenalty)
	}

	return validSubarrays + maxPenalty
}
