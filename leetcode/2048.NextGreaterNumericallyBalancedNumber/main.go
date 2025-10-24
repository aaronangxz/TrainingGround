package main

func nextBeautifulNumber(n int) int {
	i := n + 1
	for true {
		if isBeatiful(i) {
			return i
		}
		i++
	}
	return -1
}

func isBeatiful(a int) bool {
	arr := make([]int, 10)
	for a >= 1 {
		pop := a % 10
		arr[pop]++
		a = a / 10
	}
	if arr[0] != 0 {
		return false
	}
	for i := 0; i < 10; i++ {
		if arr[i] != 0 && arr[i] != i {
			return false
		}
	}
	return true
}
