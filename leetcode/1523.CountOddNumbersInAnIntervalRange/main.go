package main

func countOdds(low int, high int) int {
	oddsUpTo := func(x int) int {
		return (x + 1) / 2
	}

	return oddsUpTo(high) - oddsUpTo(low-1)
}
