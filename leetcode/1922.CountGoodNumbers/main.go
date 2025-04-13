package main

const mod = 1e9 + 7

func countGoodNumbers(n int64) int {
	// For even indices (0-based): 5 options (0,2,4,6,8)
	// For odd indices: 4 options (2,3,5,7)
	// Total = 5^(count of even indices) * 4^(count of odd indices)

	evenCount := (n + 1) / 2 // number of even indices (0, 2, 4...)
	oddCount := n / 2        // number of odd indices (1, 3, 5...)

	// Calculate (5^evenCount * 4^oddCount) % mod using fast exponentiation
	fivePow := fastExp(5, evenCount)
	fourPow := fastExp(4, oddCount)

	return (fivePow * fourPow) % mod
}

func fastExp(base, exponent int64) int {
	res := 1
	b := int(base)
	e := exponent
	b %= mod

	for e > 0 {
		if e%2 == 1 {
			res = (res * b) % mod
		}
		b = (b * b) % mod
		e /= 2
	}

	return res
}
