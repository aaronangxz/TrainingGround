package main

func largestPathValue(colors string, edges [][]int) int {
	n := len(colors)
	adj := make([][]int, n)
	freqs := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]int, 0)
		freqs[i] = make([]int, 26)
	}
	indegree := make([]int, n)
	visited := make([]bool, n)

	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		indegree[edge[1]]++
	}

	ans := -1
	q := make([]int, 0)
	for i, ind := range indegree {
		if ind == 0 {
			q = append(q, i)
			freqs[i][int(colors[i]-'a')] = 1
			visited[i] = true
		}
	}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		for _, next := range adj[curr] {
			nextColor := int(colors[next] - 'a')
			for i := 0; i < 26; i++ {
				newVal := freqs[curr][i]
				if i == nextColor {
					newVal++
				}
				freqs[next][i] = max(freqs[next][i], newVal)
			}
			indegree[next]--
			if indegree[next] == 0 {
				q = append(q, next)
				visited[next] = true
			}
		}

		for i := 0; i < 26; i++ {
			ans = max(ans, freqs[curr][i])
		}
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			return -1
		}
	}

	return ans
}
