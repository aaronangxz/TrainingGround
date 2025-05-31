package main

func snakesAndLadders(board [][]int) int {
	n := len(board)
	dp := make([]int, n*n+1)
	move := make([]int, n*n+1)
	for i := 1; i <= n*n; i++ {
		dp[i] = n * n
	}

	//Init move
	count, cur := 0, 1
	for i := n - 1; i >= 0; i-- {
		if count == 0 {
			for count < n {
				move[cur] = board[i][count]
				cur++
				count++
			}
		} else {
			for count > 0 {
				count--
				move[cur] = board[i][count]
				cur++
			}
		}
	}

	//BFS Queue
	q := make([]int, 0)
	q = append(q, 1)
	dp[1] = 0
	for len(q) > 0 {
		i := q[0]
		q = q[1:]
		for j := i + 1; j <= min(i+6, n*n); j++ {
			if move[j] > 0 {
				if dp[move[j]] > dp[i]+1 {
					q = append(q, move[j])
					dp[move[j]] = dp[i] + 1
				}
				continue
			}
			if dp[j] > dp[i]+1 {
				q = append(q, j)
				dp[j] = dp[i] + 1
			}

		}
	}

	if dp[n*n] == n*n {
		return -1
	}
	return dp[n*n]
}
