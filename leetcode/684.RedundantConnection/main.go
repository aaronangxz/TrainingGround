package main

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	parent := make([]int, n+1)
	rank := make([]int, n+1)

	for i := 1; i <= n; i++ {
		parent[i] = i
	}

	var find func(i int) int
	find = func(i int) int {
		if parent[i] != i {
			parent[i] = find(parent[i])
		}
		return parent[i]
	}

	var join func(u, v int)
	join = func(u, v int) {
		rootU := find(u)
		rootV := find(v)

		if rootU != rootV {
			if rank[rootU] > rank[rootV] {
				parent[rootV] = rootU
			} else if rank[rootU] < rank[rootV] {
				parent[rootU] = rootV
			} else {
				parent[rootV] = rootU
				rank[rootU]++
			}
		}
	}

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		if find(u) == find(v) {
			return edge
		}
		join(u, v)
	}

	return []int{}
}
