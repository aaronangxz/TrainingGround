package main

func coinChange(coins []int, amount int) int {
	// Mark default as some invalid amount, e.g. amount + 1 (impossible)
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	// Bottom up, from leaf node to root
	for a := 1; a < amount+1; a++ {
		// Test with coins
		for _, c := range coins {
			// Only if the difference between amount and coin amount is at least 0
			// Else, means coin amount is not small enough to form amount
			// i.e. coins = [10], amount = 9
			if a-c >= 0 {
				// Min solution for current amount is
				// 1 (self) + solution for the difference
				// i.e. coin = 5, amount = 8
				// dp[8] ($8) = 1 (1 $5 coin) + dp[3] (whatever coin needed to form $3)
				dp[a] = min(dp[a], 1+dp[a-c])
			}
		}
	}
	if dp[amount] != amount+1 {
		return dp[amount]
	}
	return -1
}

// Top-Down
//func coinChange(coins []int, amount int) int {
//	memo := make([]*int, amount+1)
//	var dfs func(remain int) int
//	dfs = func(remain int) int {
//		if memo[remain] != nil { return *memo[remain] }
//		if remain == 0 { return 0 }
//		if remain < 0 { return math.MaxInt32 }
//		numberOfCoins := math.MaxInt32
//		for _, c := range coins {
//			if c > remain { continue }
//			numberOfCoins = Min(numberOfCoins, 1+dfs(remain-c))
//		}
//		memo[remain] = &numberOfCoins; return numberOfCoins
//	}
//	numberOfCoins := dfs(amount)
//	if numberOfCoins == math.MaxInt32 { return -1 }
//	return numberOfCoins
//}
//
//func Min(a, b int) int { if a < b { return a }; return b }

//2D Bottom Up
//func coinChange(coins []int, amount int) int {
//	dp := make([][]int, len(coins)+1); for i := range dp { dp[i] = make([]int, amount+1) }
//	for i := 1; i <= amount; i++ { dp[0][i] = math.MaxInt32 }
//
//	for c := 1; c <= len(coins); c++ {
//		for a := 1; a <= amount; a++ {
//			if a >= coins[c-1] {
//				dp[c][a] = Min(dp[c-1][a], 1+dp[c][a-coins[c-1]])
//			} else {
//				dp[c][a] = dp[c-1][a]
//			}
//		}
//	}
//
//	if dp[len(coins)][amount] == math.MaxInt32 { return -1 }
//	return dp[len(coins)][amount]
//}
//
//func Min(a, b int) int { if a < b { return a }; return b }
