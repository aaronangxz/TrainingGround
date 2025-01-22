package main

type Coordinates struct {
	X int
	Y int
}

func highestPeak(isWater [][]int) [][]int {
	R, C := len(isWater), len(isWater[0])

	// Fill all res with -1 first
	res := make([][]int, R)
	for r := range R {
		res[r] = make([]int, C)
		for c := range C {
			res[r][c] = -1
		}
	}

	var q []Coordinates

	// Fill water cells as 0 first
	for r := range isWater {
		for c := range isWater[r] {
			if isWater[r][c] == 1 {
				q = append(q, Coordinates{X: r, Y: c})
				res[r][c] = 0
			}
		}
	}

	// BFS
	for len(q) != 0 {
		// Pop Left from queue
		curr := q[0]
		q = q[1:]

		// Height of current cell
		currHeight := res[curr.X][curr.Y]
		// Top, Down, Left, Right from curr cell
		neighbours := []Coordinates{{X: curr.X - 1, Y: curr.Y}, {X: curr.X + 1, Y: curr.Y}, {X: curr.X, Y: curr.Y + 1}, {X: curr.X, Y: curr.Y - 1}}

		for _, n := range neighbours {
			// Skip if out of bounds or if it's filled before
			if n.X < 0 || n.X == R || n.Y < 0 || n.Y == C || res[n.X][n.Y] != -1 {
				continue
			}
			// Put in queue
			q = append(q, n)
			// Increment height from current cell
			res[n.X][n.Y] = currHeight + 1
		}
	}
	return res
}
