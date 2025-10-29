package main

import "math"

func smallestNumber(n int) int {
	cnt := 0 // count needed bits
	for n > 0 {
		n /= 2
		cnt++
	}
	return int(math.Pow(2.0, float64(cnt))) - 1
}
