package main

import "math/bits"

func kthCharacter(k int64, operations []int) byte {
	mask := uint64(0)
	for i := len(operations) - 1; i >= 0; i-- {
		mask = (mask << 1) | uint64(operations[i])
	}
	return byte('a' + bits.OnesCount64((uint64(k)-1)&mask)%26)
}
