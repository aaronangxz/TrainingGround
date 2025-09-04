package main

import "math"

func findClosest(x int, y int, z int) int {
	first := int(math.Abs(float64(z - x)))
	second := int(math.Abs(float64(z - y)))
	if first == second {
		return 0
	}
	if first < second {
		return 1
	}
	return 2
}
