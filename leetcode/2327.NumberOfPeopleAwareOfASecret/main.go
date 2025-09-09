package main

func peopleAwareOfSecret(n int, delay int, forget int) int {
	const MOD = 1000000007

	know := 1
	spread := 0

	forgetOnDay := make(map[int]int)
	beginSpreadOnDay := make(map[int]int)

	beginSpreadOnDay[delay+1] = 1
	forgetOnDay[forget+1] = 1

	for day := 2; day <= n; day++ {
		spread = (spread + beginSpreadOnDay[day] - forgetOnDay[day] + MOD) % MOD
		know = (know + spread - forgetOnDay[day] + MOD) % MOD

		beginSpreadOnDay[day+delay] = spread
		forgetOnDay[day+forget] = spread
	}

	return know
}
