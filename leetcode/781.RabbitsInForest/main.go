package main

func numRabbits(answers []int) int {
	count := map[int]int{}
	for _, a := range answers {
		count[a]++
	}

	res := 0
	for k, v := range count {
		res += ((v + k) / (k + 1)) * (k + 1)
	}
	return res
}
