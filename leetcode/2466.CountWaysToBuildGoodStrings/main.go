package main

const (
	mod = 1000000007
)

// Backtracking
//func countGoodStrings(low int, high int, zero int, one int) int {
//	dp := make(map[int]int)
//
//	var dfs func(int) int
//	dfs = func(length int) int {
//		if length > high {
//			return 0
//		}
//		if v, ok := dp[length]; ok {
//			return v
//		}
//		if length >= low {
//			dp[length] = 1
//		} else {
//			dp[length] = 0
//		}
//		dp[length] += dfs(length+zero) + dfs(length+one)
//		return dp[length] % mod
//	}
//	return dfs(0)
//}

// DP
func countGoodStrings(low int, high int, zero int, one int) int {
	dp := make(map[int]int)
	dp[0] = 1
	for i := 1; i <= high; i++ {
		dp[i] = (dp[i-one] + dp[i-zero]) % mod
	}

	sum := 0

	for k, v := range dp {
		if k >= low && k <= high {
			sum += v
		}
	}
	return sum
}
