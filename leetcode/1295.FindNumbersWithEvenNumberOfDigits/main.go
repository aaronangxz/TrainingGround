package main

func findNumbers(nums []int) int {
	count := 0
	for _, n := range nums {
		digits := 0

		for n > 0 {
			digits++
			n /= 10
		}
		if digits%2 == 0 {
			count++
		}
	}
	return count
}
