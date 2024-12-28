package main

import "github.com/aaronangxz/TrainingGround/common"

func maxScoreSightseeingPair(values []int) int {
	// Formula = values[i] + values[j] + i - j -> is basically (values[i] + i) + (values[j] - j)
	// Taking i as the 'Left' and j as the 'Right'
	// Everytime we move towards the back, calculate the current values[j] - j, which is 'Right'
	// We can view everything in prev idx as the 'Left' -> as long as we maintain the max on the 'Left'
	maxScore := 0

	// First element is fixed, since idx 0 has no 'Left', logically it's the largest 'Left'
	maxLeftScore := values[0]

	// Start from idx 1
	for i := 1; i < len(values); i++ {
		// Viewing it as 'Right' now (or j in the formula), so - i
		currRight := values[i] - i
		// As soon as we know the current 'Right', we can compare if 'Left' + 'Right' is the largest thus far
		maxScore = common.Max(maxScore, maxLeftScore+currRight)
		// Afterward, update the largest 'Left', including the current value that we are looking at
		// Viewing it as 'Left' now (or i in the formula), so + i
		currLeft := values[i] + i
		maxLeftScore = common.Max(maxLeftScore, currLeft)
	}
	return maxScore
}
