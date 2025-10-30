package main

func minNumberOperations(target []int) int {
	n := len(target)
	ans := target[0]
	for i := 1; i < n; i++ {
		if target[i] > target[i-1] {
			ans += target[i] - target[i-1]
		}
	}
	return ans
}
