package main

import "github.com/aaronangxz/TrainingGround/common"

func maxScore(cardPoints []int, k int) int {
	//sum of the window (first k elements)
	sum := 0
	for i := 0; i < k; i++ {
		sum += cardPoints[i]
	}

	//make a copy for comparison later
	curr := sum

	//slide the window to the left
	//i.e. the front of window move one index to the left, then wraparound to the back of array
	//window covers 1,2,3 : [1 2 3] 4 5
	//window covers 5,1,2 : 1 2 ] 3 4 [ 5
	//window covers 1,4,5 : 1 ] 2 3 [ 4 5
	for i := k - 1; i >= 0; i-- {
		//remove the current last index
		curr -= cardPoints[i]

		//add the next index in wraparound
		curr += cardPoints[len(cardPoints)-k+i]

		//we want the maximum possible
		sum = common.Max(sum, curr)
	}
	return sum
}
