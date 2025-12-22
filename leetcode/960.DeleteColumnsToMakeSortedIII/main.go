package main

func minDeletionSize(strs []string) int {
	deletionsToColumn := make([]int, len(strs[0])+1) // deletionsToColumn is 1 place longer, to accomodate for the very important End node
	// this loop simulates going from our placeholder start node to every other node (the cost is just the column index)
	// also there will always be a path from S to any other node, it translates to where we choose to start
	for i := 0; i < len(deletionsToColumn); i++ {
		deletionsToColumn[i] = i
	}
	return minDeletionsRoute(strs, deletionsToColumn)
}

// I guess there was no real need for me to split to another function, just a remainder of my previous DFS attempt
func minDeletionsRoute(strs []string, deletionsToColumn []int) int {
	for column := 0; column < len(strs[0]); column++ {
		for nextColumn := column + 1; nextColumn < len(deletionsToColumn); nextColumn++ {
			// it's cheaper (computationally) to check if the cost of the edge is cheaper than the previous best way to arrive to it, so that condition comes before the call to hasRoute
			if deletionsToColumn[nextColumn] >= deletionsToColumn[column]+nextColumn-column-1 && hasRoute(strs, column, nextColumn) {
				deletionsToColumn[nextColumn] = deletionsToColumn[column] + nextColumn - column - 1
			}
		}
	}
	return deletionsToColumn[len(deletionsToColumn)-1]
}

// hasRoute checks that all the chars in the column (c) are in order with the chars in the next column (nc)
func hasRoute(strs []string, c, nc int) bool {
	if c == -1 || nc == len(strs[0]) {
		return true
	} // if nc is out of bounds, that means it's the end node, to which all nodes are free to reach (at a cost!)
	for si := 0; si < len(strs); si++ {
		if strs[si][c] > strs[si][nc] {
			return false
		}
	}
	return true
}
