package main

const mod = 1e9 + 7

func numOfSubarrays(arr []int) int {
	currSum := 0
	odd := 0
	even := 0
	res := 0

	for _, num := range arr {
		currSum += num
		if currSum%2 == 0 {
			// If the current prefix sum is even, the number of subarrays with odd sum is the number of odd prefix sums so far
			res = (res + odd) % mod
			even++
		} else {
			// If the current prefix sum is odd, the number of subarrays with odd sum is the number of even prefix sums so far
			res = (res + 1 + even) % mod
			odd++
		}
	}

	return res
}
