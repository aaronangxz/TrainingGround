package main

import (
	"math"
	"strconv"
	"strings"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	num, den := int64(numerator), int64(denominator)
	sign := ""
	if (num < 0) != (den < 0) {
		sign = "-"
	}
	num, den = int64(math.Abs(float64(num))), int64(math.Abs(float64(den)))
	var b strings.Builder
	b.WriteString(sign)
	b.WriteString(strconv.FormatInt(num/den, 10))
	rem := num % den
	if rem == 0 {
		return b.String()
	}
	b.WriteByte('.')
	seen := map[int64]int{}
	for rem != 0 {
		if pos, ok := seen[rem]; ok {
			s := b.String()
			return s[:pos] + "(" + s[pos:] + ")"
		}
		seen[rem] = b.Len()
		rem *= 10
		b.WriteByte(byte('0' + rem/den))
		rem %= den
	}
	return b.String()
}
