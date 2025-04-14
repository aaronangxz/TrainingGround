package main

func countGoodTriplets(arr []int, a int, b int, c int) int {
	n := len(arr)
	res := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if calculate(arr[i], arr[j]) <= a && calculate(arr[j], arr[k]) <= b && calculate(arr[i], arr[k]) <= c {
					res++
				}
			}
		}
	}
	return res
}

func calculate(a, b int) int {
	diff := a - b
	if diff < 0 {
		return -diff
	}
	return diff
}
