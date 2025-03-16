package main

import (
	"math"
)

func repairCars(ranks []int, cars int) int64 {
	l, r := 1, ranks[0]*cars*cars
	res := -1

	for l <= r {
		m := (l + r) / 2

		repairCount := 0
		for _, r := range ranks {
			repairCount += int(math.Sqrt(float64(m / r)))
		}
		if repairCount >= cars {
			res = m
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return int64(res)
}
