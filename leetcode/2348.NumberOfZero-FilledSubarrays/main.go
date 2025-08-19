package main

func zeroFilledSubarray(nums []int) int64 {
	var ans, run int64
	for _, x := range nums {
		if x == 0 {
			run++
			ans += run
		} else {
			run = 0
		}
	}
	return ans
}
