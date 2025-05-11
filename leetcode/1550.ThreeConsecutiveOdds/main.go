package main

func threeConsecutiveOdds(arr []int) bool {
	currCount := 0
	for _, a := range arr {
		if a%2 != 0 {
			currCount++
			if currCount == 3 {
				return true
			}
		} else {
			currCount = 0
		}
	}
	return false
}
