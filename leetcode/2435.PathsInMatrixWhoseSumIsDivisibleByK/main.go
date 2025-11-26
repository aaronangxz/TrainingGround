package main

const MOD = 1e9 + 7

func numberOfPaths(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])

	dp := make([][][]int, m+1)
	for i := range m + 1 {
		dp[i] = make([][]int, n+1)
		for j := range n + 1 {
			dp[i][j] = make([]int, k)
		}
	}
	dp[0][1][0] = 1

	for r := range m {
		for c := range n {
			num := grid[r][c]
			for rem := range k {
				if dp[r+1][c][rem] > 0 {
					dp[r+1][c+1][(rem+num)%k] += dp[r+1][c][rem]
					dp[r+1][c+1][(rem+num)%k] %= MOD
				}
				if dp[r][c+1][rem] > 0 {
					dp[r+1][c+1][(rem+num)%k] += dp[r][c+1][rem]
					dp[r+1][c+1][(rem+num)%k] %= MOD
				}
			}
		}
	}

	return dp[m][n][0]
}
