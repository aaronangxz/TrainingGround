package main

func numWaterBottles(numBottles int, numExchange int) int {
	total := numBottles

	for numBottles >= numExchange {
		curr := numBottles / numExchange
		rem := numBottles % numExchange
		total += curr
		numBottles = rem + numBottles/numExchange
	}

	return total
}
