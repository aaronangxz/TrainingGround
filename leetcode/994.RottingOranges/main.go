package main

type Cor struct {
	row int
	col int
}

var directions = []Cor{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func orangesRotting(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	mins := 0
	fresh := 0
	var q []Cor

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				fresh++
			} else if grid[i][j] == 2 {
				q = append(q, Cor{i, j})
			}
		}
	}

	for len(q) > 0 && fresh > 0 {
		size := len(q)

		for i := 0; i < size; i++ {
			curr := q[0]
			q = q[1:]

			for _, d := range directions {
				newR, newC := curr.row+d.row, curr.col+d.col
				if newR < 0 || newR >= rows || newC < 0 || newC >= cols || grid[newR][newC] != 1 {
					continue
				}
				grid[newR][newC] = 2
				q = append(q, Cor{newR, newC})
				fresh--
			}
		}
		mins++
	}
	if fresh != 0 {
		return -1
	}
	return mins
}
