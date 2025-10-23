package main

func hasSameDigits(s string) bool {
	digits := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		digits[i] = int(s[i] - '0')
	}

	for len(digits) > 2 {
		next := make([]int, len(digits)-1)
		for i := 0; i+1 < len(digits); i++ {
			next[i] = (digits[i] + digits[i+1]) % 10
		}
		digits = next
	}

	return len(digits) == 2 && digits[0] == digits[1]
}
