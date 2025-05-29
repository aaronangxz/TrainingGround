package main

func getColors(edges [][]int) []int {
	n := len(edges) + 1
	adj := make([][]int, n)
	colors := make([]int, n)
	for i := 0; i < n; i++ {
		colors[i] = -1
	}
	for _, edge := range edges {
		u := edge[0]
		v := edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	q := []int{0}
	colors[0] = 0
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		for _, next := range adj[curr] {
			if colors[next] == -1 {
				colors[next] = colors[curr] ^ 1
				q = append(q, next)
			}
		}
	}

	return colors
}

func maxTargetNodes(edges1 [][]int, edges2 [][]int) []int {
	n := len(edges1) + 1
	color1 := getColors(edges1)
	color2 := getColors(edges2)
	white1, black1, white2, black2 := 0, 0, 0, 0

	for _, col := range color1 {
		if col == 0 {
			white1++
		} else {
			black1++
		}
	}
	for _, col := range color2 {
		if col == 0 {
			white2++
		} else {
			black2++
		}
	}
	max2 := max(white2, black2)

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		if color1[i] == 0 {
			ans[i] = white1 + max2
		} else {
			ans[i] = black1 + max2
		}
	}

	return ans
}
