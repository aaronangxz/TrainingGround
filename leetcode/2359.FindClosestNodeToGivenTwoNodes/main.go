package main

func getDistances(edges []int, start int) []int {
	n := len(edges)
	dist := make([]int, n)
	for i := range dist {
		dist[i] = -1
	}
	d := 0
	for start != -1 && dist[start] == -1 {
		dist[start] = d
		d++
		start = edges[start]
	}
	return dist
}

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	dist1 := getDistances(edges, node1)
	dist2 := getDistances(edges, node2)
	result := -1
	minDistance := int(^uint(0) >> 1)

	for i := 0; i < len(edges); i++ {
		if dist1[i] != -1 && dist2[i] != -1 {
			maxDist := dist1[i]
			if dist2[i] > maxDist {
				maxDist = dist2[i]
			}
			if maxDist < minDistance {
				minDistance = maxDist
				result = i
			}
		}
	}
	return result
}
