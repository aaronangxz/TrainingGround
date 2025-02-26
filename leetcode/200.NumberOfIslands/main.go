package main

type Cor struct {
	row int
	col int
}

var directions = []Cor{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func numIslands(grid [][]byte) int {
	row, col := len(grid), len(grid[0])
	visited := make([][]bool, row)
	for i := range visited {
		visited[i] = make([]bool, col)
	}
	islands := 0

	var bfs func(r, c int)
	bfs = func(r, c int) {
		var q []Cor
		visited[r][c] = true
		q = append(q, Cor{r, c})

		for len(q) > 0 {
			curr := q[0]
			q = q[1:]

			for _, d := range directions {
				newR, newC := curr.row+d.row, curr.col+d.col
				if newR >= 0 && newR < row && newC >= 0 && newC < col && grid[newR][newC] == '1' && !visited[newR][newC] {
					q = append(q, Cor{newR, newC})
					visited[newR][newC] = true
				}
			}
		}

	}

	for r := range row {
		for c := range col {
			if grid[r][c] == '1' && !visited[r][c] {
				bfs(r, c)
				islands++
			}
		}
	}
	return islands
}
