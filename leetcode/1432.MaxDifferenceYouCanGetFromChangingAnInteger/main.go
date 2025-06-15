package main

import "strconv"

func maxDiff(num int) int {
	min, max, str := num, num, strconv.Itoa(num)
	// find min/max by performing all possible remap operations
	for i := range 10 {
		for j := range 10 {
			n := remap(str, i, j)
			if n > 0 && len(strconv.Itoa(n)) == len(str) {
				if n < min {
					min = n
				}
				if n > max {
					max = n
				}
			}
		}
	}
	return max - min
}

// remap i to j in a number represented as a string
func remap(str string, i int, j int) int {
	n := 0
	for p := range len(str) {
		d := int(str[p] - '0')
		if d == i {
			n = n*10 + j
		} else {
			n = n*10 + d
		}
	}
	return n
}
