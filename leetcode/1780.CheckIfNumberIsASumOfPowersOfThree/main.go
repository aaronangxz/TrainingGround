package main

import (
	"math"
)

func checkPowersOfThree(n int) bool {
	// Find the ceiling of power
	maxI := 0
	nC := n
	for nC >= 0 {
		pow := int(math.Pow(3, float64(maxI)))
		maxI++
		nC -= pow
	}

	// Starting from the largest power, pick whichever that does not exceed the remaining n
	for maxI >= 0 {
		pow := int(math.Pow(3, float64(maxI)))
		if pow <= n {
			n -= pow
		}
		maxI--
	}

	// If at the end n is not fully used, it is not a sum of powers of 3
	return n == 0
}
