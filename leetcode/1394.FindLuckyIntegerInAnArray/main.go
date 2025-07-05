package main

func findLucky(arr []int) int {
	cnt := make([]int, 501)

	for _, val := range arr {
		cnt[val]++
	}

	for i := 500; i >= 1; i-- {
		if i == cnt[i] {
			return i
		}
	}
	return -1
}
