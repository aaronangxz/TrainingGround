package main

func minOperations(boxes string) []int {
	ans := make([]int, len(boxes))

	// Keeping track of the total number of balls seen in boxes
	// and the moves required to move to ith index
	// Forward pass (moving balls from left to right)
	balls, moves := 0, 0
	for i := range boxes {
		// store the sum in ans slice
		ans[i] = balls + moves
		// the number of balls so far is basically the move needed
		moves += balls
		// if the current box is not empty, increment ball count
		if string(boxes[i]) == "1" {
			balls++
		}
	}
	// Backward pass (moving balls from right to left)
	balls, moves = 0, 0
	for i := len(boxes) - 1; i >= 0; i-- {
		// here, we need to increment whatever gotten from the left-right pass
		ans[i] += balls + moves
		moves += balls
		if string(boxes[i]) == "1" {
			balls++
		}
	}
	return ans
}
