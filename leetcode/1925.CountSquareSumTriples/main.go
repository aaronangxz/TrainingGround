package main

// Complexity:
// Time O(nLog(n)) and Space O(1).
func countTriples(n int) int {
	// Euclid's Formula
	result := 0
	for i := 1; ; i++ {
		ii := i * i
		if ii*2 >= n {
			break
		}

		for j := i + 1; ; j += 2 {
			c := ii + j*j
			if c > n {
				break
			}

			if gcd(i, j) == 1 {
				result += n / c
			}
		}
	}
	return result * 2 // We can swap a and b in the triples
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
