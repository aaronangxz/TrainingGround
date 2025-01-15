package main

import "math/bits"

func minimizeXor(num1 int, num2 int) int {
	var ans int

	ptr, bits := bits.Len(uint(num1)), bits.OnesCount(uint(num2))
	if bits > ptr {
		ptr = bits
	}

	for ; ptr >= 0 && bits > 0; ptr-- {
		if num1>>ptr&1 == 1 || ptr < bits {
			ans |= 1 << ptr
			bits--
		}
	}

	return ans
}
