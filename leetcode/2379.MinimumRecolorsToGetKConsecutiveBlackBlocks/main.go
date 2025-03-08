package main

import "math"

func minimumRecolors(blocks string, k int) int {
	i := 0
	j := k
	minTimes := math.MaxInt32

	for j <= len(blocks) {
		times := 0
		for idx := i; idx < j; idx++ {
			if string(blocks[idx]) == "W" {
				times++
			}
		}
		minTimes = min(minTimes, times)
		i++
		j++
	}
	return minTimes
}
