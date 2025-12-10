package main

func countPermutations(complexity []int) int {
	MOD := int(1e9 + 7)
	for i := 1; i < len(complexity); i++ {
		if complexity[i] <= complexity[0] {
			return 0
		}
	}

	res := 1
	for i := 1; i < len(complexity); i++ {
		res *= i
		res %= MOD
	}
	return res
}
