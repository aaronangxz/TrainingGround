package main

func smallestRepunitDivByK(K int) int {
	for n, i := 0, 1; i <= K; i++ {
		n = n*10 + 1
		n = n % K
		if n == 0 {
			return i
		}
	}
	return -1
}
