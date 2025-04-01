package main

func mostPoints(questions [][]int) int64 {
	n := len(questions)
	dp := make([]int64, n+1) // dp[i] represents max points from question i to end

	for i := n - 1; i >= 0; i-- {
		points, brainpower := questions[i][0], questions[i][1]
		next := i + brainpower + 1

		// If we solve current question, we can't solve next 'brainpower' questions
		solve := int64(points)
		if next < n {
			solve += dp[next]
		}

		// If we skip current question, we get dp[i+1]
		skip := dp[i+1]

		// Choose the better option
		if solve > skip {
			dp[i] = solve
		} else {
			dp[i] = skip
		}
	}

	return dp[0]
}
