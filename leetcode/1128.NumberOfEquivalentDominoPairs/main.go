package main

func numEquivDominoPairs(dominoes [][]int) int {
	count := 0
	// Constraints: 1 <= dominoes[i][j] <= 9, max number is 99
	dom := make([]int, 100)
	for _, a := range dominoes {
		// lowest digit in front
		if a[1] < a[0] {
			a[0], a[1] = a[1], a[0]
		}
		num := a[0]*10 + a[1]
		count += dom[num]
		dom[num] += 1
	}

	return count
}
