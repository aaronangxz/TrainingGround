package main

func lenOfVDiagonal(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	directions := []int{1, 1, -1, -1, 1}
	cache := make([][][4][2]int, rows)
	for i := range cache {
		cache[i] = make([][4][2]int, cols)
	}

	var explore func(r, c, dir, turn int) int
	explore = func(r, c, dir, turn int) int {
		if cache[r][c][dir][turn] != 0 {
			return cache[r][c][dir][turn]
		}

		nr := r + directions[dir]
		nc := c + directions[dir+1]

		var expect int
		if grid[r][c] == 1 {
			expect = 2
		} else {
			expect = 2 - grid[r][c]
		}

		if nr < 0 || nr >= rows || nc < 0 || nc >= cols || grid[nr][nc] != expect {
			cache[r][c][dir][turn] = 0
			return 0
		}

		length := explore(nr, nc, dir, turn)
		if turn > 0 {
			length = max(length, explore(nr, nc, (dir+1)%4, 0))
		}
		cache[r][c][dir][turn] = length + 1
		return length + 1
	}

	result := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 {
				for dir := 0; dir < 4; dir++ {
					result = max(result, explore(r, c, dir, 1)+1)
				}
			}
		}
	}
	return result
}
