package main

const MOD int = 1e9 + 7

func modInverse(a, mod int) int {
	res := 1
	power := mod - 2
	for power > 0 {
		if power&1 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		power >>= 1
	}
	return res
}

func nCr(n, r int) int {
	if r > n {
		return 0
	}
	if r == 0 || r == n {
		return 1
	}

	res := 1
	for i := 1; i <= r; i++ {
		res = res * (n - r + i) % MOD
		res = res * modInverse(i, MOD) % MOD
	}

	return res
}

func binExpo(a, b int) int {
	result := 1
	for b > 0 {
		if b&1 == 1 {
			result = result * a % MOD
		}
		a = a * a % MOD
		b >>= 1
	}
	return result
}

func countGoodArrays(n int, m int, k int) int {
	formula := m * binExpo(m-1, n-(k+1)) % MOD
	updatedFormula := formula * nCr(n-1, k) % MOD
	return updatedFormula
}
