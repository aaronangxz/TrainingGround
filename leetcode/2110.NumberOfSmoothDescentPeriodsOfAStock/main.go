package main

func getDescentPeriods(prices []int) int64 {
	var ans int64 = 1
	var length int64 = 1

	for i := 1; i < len(prices); i++ {
		if prices[i] == prices[i-1]-1 {
			length++
		} else {
			length = 1
		}
		ans += length
	}
	return ans
}
