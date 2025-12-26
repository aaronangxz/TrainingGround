package main

func bestClosingTime(customers string) int {
	var n int = len(customers)
	var pen = 0
	var min = 0
	var result = n
	for indx := n - 1; indx >= 0; indx-- {
		if customers[indx] == 'Y' {
			pen++
		} else {
			pen--
		}
		if pen <= min {
			min = pen
			result = indx
		}
	}
	return result
}
