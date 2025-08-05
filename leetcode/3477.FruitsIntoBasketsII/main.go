package main

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	n := len(fruits)
	alloted := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if fruits[i] <= baskets[j] {
				baskets[j] = -1 // mark as used
				alloted++
				break
			}
		}
	}
	return n - alloted
}
