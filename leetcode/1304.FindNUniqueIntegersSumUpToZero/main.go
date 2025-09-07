package main

func sumZero(n int) []int {
	arr := make([]int, n)
	k := 1

	for i := 0; i < n/2; i++ {
		arr[i] = k
		arr[n-1-i] = -k
		k++
	}
	return arr
}
