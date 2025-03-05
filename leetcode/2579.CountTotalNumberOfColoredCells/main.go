package main

func coloredCells(n int) int64 {
	var total int64 = 1
	n--
	for n > 0 {
		total += int64(4 * n)
		n--
	}
	return total
}
