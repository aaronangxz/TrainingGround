package main

func divideArray(nums []int) bool {
	m := make(map[int]int)

	for _, n := range nums {
		m[n]++
	}

	for _, mm := range m {
		if mm%2 != 0 {
			return false
		}
	}
	return true
}
