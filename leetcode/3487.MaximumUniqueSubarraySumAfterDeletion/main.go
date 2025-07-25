package main

import "math"

func maxSum(nums []int) int {
	mn := math.MinInt32
	seen := make(map[int]bool)
	sum := 0

	for _, val := range nums {
		if !seen[val] {
			// if element first occurured, hence unique
			if val >= 0 {
				sum += val
			} else {
				if val > mn {
					mn = val
				}
			}
		}

		seen[val] = true
	}

	if sum == 0 && !seen[0] {
		return mn
	}

	return sum
}
