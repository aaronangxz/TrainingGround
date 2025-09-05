package main

func makeTheIntegerZero(n1 int, n2 int) int {
	for op := 1; op <= 60; op++ {
		rem := int64(n1) - int64(op)*int64(n2)
		if rem < int64(op) {
			continue
		}
		pop := 0
		tmp := rem
		for tmp > 0 {
			pop += int(tmp & 1)
			tmp >>= 1
		}
		if pop <= op && int64(op) <= rem {
			return op
		}
	}
	return -1
}
