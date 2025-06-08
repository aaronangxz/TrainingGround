package main

func lexicalOrder(n int) []int {
	var ans []int

	var solve func(i int)
	solve = func(i int) {
		// base conditon - where recursion gets over
		if i > n {
			return
		}

		cur := i // 1
		ans = append(ans, cur)
		for j := 0; j <= 9; j++ {
			tmp := cur*10 + j
			solve(tmp)
		}
	}

	for i := 1; i <= 9; i++ {
		solve(i)
	}

	return ans
}
