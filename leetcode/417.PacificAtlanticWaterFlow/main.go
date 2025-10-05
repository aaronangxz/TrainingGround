package main

func pacificAtlantic(heights [][]int) [][]int {
	res := [][]int{}

	ROW, COL := len(heights), len(heights[0])

	pacific := make([][]bool, ROW)
	atlantic := make([][]bool, ROW)
	for r := 0; r < ROW; r++ {
		pacific[r] = make([]bool, COL)
		atlantic[r] = make([]bool, COL)
	}

	var dfs func(r int, c int, visit [][]bool, prevHeight int)
	dfs = func(r int, c int, visit [][]bool, prevHeight int) {
		if r < 0 || c < 0 || r == ROW || c == COL || heights[r][c] < prevHeight || visit[r][c] {
			return
		}

		visit[r][c] = true
		dfs(r+1, c, visit, heights[r][c])
		dfs(r-1, c, visit, heights[r][c])
		dfs(r, c+1, visit, heights[r][c])
		dfs(r, c-1, visit, heights[r][c])
	}

	for c := 0; c < COL; c++ {
		dfs(0, c, pacific, heights[0][c])
		dfs(ROW-1, c, atlantic, heights[ROW-1][c])
	}
	for r := 0; r < ROW; r++ {
		dfs(r, 0, pacific, heights[r][0])
		dfs(r, COL-1, atlantic, heights[r][COL-1])
	}

	for r := 0; r < ROW; r++ {
		for c := 0; c < COL; c++ {
			if pacific[r][c] && atlantic[r][c] {
				res = append(res, []int{r, c})
			}
		}
	}
	return res
}
