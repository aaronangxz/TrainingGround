package main

const mod = 1000000007

func multiply(a, b [26][26]int64) (c [26][26]int64) {
	for i := 0; i < 26; i++ {
		for k := 0; k < 26; k++ {
			if aik := a[i][k]; aik != 0 {
				for j := 0; j < 26; j++ {
					c[i][j] = (c[i][j] + aik*b[k][j]) % mod
				}
			}
		}
	}
	return
}

func matrixPower(m [26][26]int64, exp int) (res [26][26]int64) {
	for i := 0; i < 26; i++ {
		res[i][i] = 1
	}

	base := m
	for exp > 0 {
		if exp&1 == 1 {
			res = multiply(res, base)
		}

		base = multiply(base, base)
		exp >>= 1
	}

	return
}

func lengthAfterTransformations(s string, t int, nums []int) int {
	var freq [26]int64
	for _, ch := range s {
		freq[ch-'a']++
	}

	var baseM [26][26]int64
	for i := 0; i < 26; i++ {
		for k := 1; k <= nums[i]; k++ {
			j := (i + k) % 26
			baseM[i][j]++
		}
	}

	mt := matrixPower(baseM, t)
	var result int64
	for i := 0; i < 26; i++ {
		if fi := freq[i]; fi != 0 {
			for j := 0; j < 26; j++ {
				result = (result + fi*mt[i][j]) % mod
			}
		}
	}

	return int(result)
}
