package main

func specialTriplets(nums []int) int {
	const MOD int64 = 1_000_000_007

	right := make(map[int]int64)
	left := make(map[int]int64)

	for _, x := range nums {
		right[x]++
	}

	var ans int64 = 0

	for _, x := range nums {
		right[x]--

		target := x * 2

		cntLeft := left[target]
		cntRight := right[target]

		add := (cntLeft * cntRight) % MOD
		ans = (ans + add) % MOD

		left[x]++
	}

	return int(ans % MOD)
}
