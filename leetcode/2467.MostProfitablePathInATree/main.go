package main

import "math"

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	graph := make([][]int32, len(amount))
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], int32(edge[1]))
		graph[edge[1]] = append(graph[edge[1]], int32(edge[0]))
	}
	parents := make([]int32, len(amount))
	for i := range parents {
		parents[i] = -1
	}
	parents[0] = -2
	var fillParents func(parentIdx int32)
	fillParents = func(parentIdx int32) {
		for _, childIdx := range graph[parentIdx] {
			if parents[childIdx] == -1 {
				parents[childIdx] = parentIdx
				fillParents(childIdx)
			}
		}
	}
	fillParents(0)

	slow, fast := int32(bob), parents[bob]
	for fast != 0 && parents[fast] != 0 {
		amount[slow] = 0
		slow = parents[slow]
		fast = parents[parents[fast]]
	}
	amount[slow] = 0
	if fast != 0 {
		amount[parents[slow]] >>= 1
	}

	var isLeaf bool
	result := math.MinInt
	for cur, next := []int32{0}, []int32{}; len(cur) != 0; cur, next = next, cur[:0] {
		for _, parent := range cur {
			isLeaf = true
			for _, child := range graph[parent] {
				if parents[child] != parent {
					continue
				}
				isLeaf = false
				amount[child] += amount[parent]
				next = append(next, child)
			}
			if isLeaf {
				result = max(result, amount[parent])
			}
		}
	}
	return result
}
