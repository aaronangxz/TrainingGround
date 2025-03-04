package main

func pivotArray(nums []int, pivot int) []int {
	var less, p, more []int
	for _, n := range nums {
		if n < pivot {
			less = append(less, n)
		} else if n > pivot {
			more = append(more, n)
		} else {
			p = append(p, n)
		}
	}
	less = append(less, p...)
	less = append(less, more...)
	return less
}
