package main

import "math"

func minScoreTriangulation(values []int) int {

	n := len(values)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	return scoreMem(values, 0, n-1, dp)
}

func scoreMem(values []int, i, j int, dp [][]int) int {
	if i+1 == j {
		return 0
	}

	if dp[i][j] != 0 {
		return dp[i][j]
	}
	ans := math.MaxInt
	for k := i + 1; k < j; k++ {

		ans = min(ans, values[i]*values[j]*values[k]+scoreMem(values, i, k, dp)+scoreMem(values, k, j, dp))
	}
	dp[i][j] = ans
	return dp[i][j]
}
