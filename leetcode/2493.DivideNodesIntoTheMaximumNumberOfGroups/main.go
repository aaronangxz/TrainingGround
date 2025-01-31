package main

func magnificentSets(n int, edges [][]int) int {
	graph := make([][]int16, n)
	for _, edge := range edges {
		graph[edge[0]-1] = append(graph[edge[0]-1], int16(edge[1]-1))
		graph[edge[1]-1] = append(graph[edge[1]-1], int16(edge[0]-1))
	}
	visited := make([]int16, n)
	var components [][]int16
	var depth int16
	var cur, next []int16
	for node := range graph {
		if visited[node] != 0 {
			continue
		}
		visited[node] = depth + 1
		component := []int16{int16(node)}
		for cur = append(cur, int16(node)); len(cur) != 0; cur, next = next, cur[:0] {
			depth++
			for _, node := range cur {
				for _, nextNode := range graph[node] {
					if visited[nextNode] != 0 {
						if visited[nextNode] != depth-1 && visited[nextNode] != depth+1 {
							return -1
						}
						continue
					}
					visited[nextNode] = depth + 1
					component = append(component, nextNode)
					next = append(next, nextNode)
				}
			}
		}
		components = append(components, component)
	}
	clear(visited)
	var result, maxComponentDepth int16
	for _, component := range components {
		maxComponentDepth = 0
		for _, node := range component {
			depth = 0
			visited[node] = 1
			for cur = append(cur, node); len(cur) != 0; cur, next = next, cur[:0] {
				depth++
				for _, node := range cur {
					for _, nextNode := range graph[node] {
						if visited[nextNode] != 0 {
							continue
						}
						visited[nextNode] = 1
						next = append(next, nextNode)
					}
				}
			}
			maxComponentDepth = max(maxComponentDepth, depth)
			for _, node = range component {
				visited[node] = 0
			}
		}
		result += maxComponentDepth
	}
	return int(result)
}
