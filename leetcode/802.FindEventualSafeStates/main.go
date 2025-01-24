package main

func eventualSafeNodes(graph [][]int) []int {
	safeNodes := make(map[int]bool)
	var res []int

	// DFS to recursively check the linked vertices
	var dfs func(idx int) bool
	dfs = func(idx int) bool {
		// If it has been checked before, directly return result
		if v, ok := safeNodes[idx]; ok {
			return v
		}

		// Else, mark it as unsafe first
		safeNodes[idx] = false

		// Check all of the linked nodes
		for _, n := range graph[idx] {
			// As long as it hits one fail, return as unsafe node
			if !dfs(n) {
				return false
			}
		}
		// Else this is a safe node
		safeNodes[idx] = true
		return true
	}

	// Ensure result is in asc order of idx
	for i := range graph {
		if dfs(i) {
			res = append(res, i)
		}
	}

	return res
}
