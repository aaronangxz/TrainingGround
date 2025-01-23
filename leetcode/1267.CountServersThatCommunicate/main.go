package main

func countServers(grid [][]int) int {
	R, C := len(grid), len(grid[0])

	// Calculate the total servers of each rows and columns
	rMap, cMap := make([]int, R), make([]int, C)

	for r := range R {
		for c := range C {
			if grid[r][c] == 1 {
				rMap[r]++
				cMap[c]++
			}
		}
	}

	total := 0

	for r := range R {
		for c := range C {
			if grid[r][c] == 1 {
				// As long as the row and column the server is on is not 1
				// It is connected
				if rMap[r] == 1 && cMap[c] == 1 {
					continue
				}
				total++
			}

		}
	}

	return total
}

// r [2,1,1,1]
// c [1,1,2,1]

// (0,0) -> r[0] == 2 & c[0] == 1 --> connected +1
// (0,1) -> r[0] == 2 & c[1] == 1 --> connected +1
// (1,2) -> r[1] == 1 & c[2] == 2 --> connected +1
// (2,2) -> r[2] == 1 & c[2] == 2 --> connected +1
// (3,4) -> r[3] == 1 & c[4] == 1 --> x, not counted
