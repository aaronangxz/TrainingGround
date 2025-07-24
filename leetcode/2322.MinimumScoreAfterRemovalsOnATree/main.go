package main

import (
	"math"
	"sort"
)

func minimumScore(nums []int, edges [][]int) int {

	n := len(nums)

	// store connections for each node
	connections := make([][]int, n)
	for _, edge := range edges {
		i, j := edge[0], edge[1]
		connections[i] = append(connections[i], j)
		connections[j] = append(connections[j], i)
	}

	// we will use enterTime and exitTime to find if one node is a parent of the another
	enterTime, exitTime := make([]int, n), make([]int, n)
	// xor of all the values of the subtree with given root
	subTreeXor := make([]int, n)
	// fill in enterTime, exitTime and subTreeXor
	var visit func(node, parent int)
	time := 0
	visit = func(node, parent int) {
		enterTime[node] = time
		time++
		xor := nums[node]
		for _, child := range connections[node] {
			if child == parent {
				continue
			}
			visit(child, node)
			xor ^= subTreeXor[child]
		}
		subTreeXor[node] = xor
		exitTime[node] = time
	}
	visit(0, -1)

	// returns true if parent is parent of the child
	isParent := func(parent, child int) bool {
		childEnterTime := enterTime[child]
		return childEnterTime > enterTime[parent] && childEnterTime < exitTime[parent]
	}

	minDiff := math.MaxInt

	// first and second are roots of the subtrees we decide to disconnect
	for first := 1; first < n; first++ {
		for second := first + 1; second < n; second++ {
			// find xors of each segment resulting from disconnection
			segmentXors := [3]int{}
			if isParent(first, second) {
				// first is parent of second:
				// 1 segment: root without first
				// 2 segment: first without second
				// 3 segment: second
				segmentXors[0] = subTreeXor[0] ^ subTreeXor[first]
				segmentXors[1] = subTreeXor[first] ^ subTreeXor[second]
				segmentXors[2] = subTreeXor[second]
			} else if isParent(second, first) {
				// second is parent of first:
				// 1 segment: root without second
				// 2 segment: second without first
				// 3 segment: first
				segmentXors[0] = subTreeXor[0] ^ subTreeXor[second]
				segmentXors[1] = subTreeXor[second] ^ subTreeXor[first]
				segmentXors[2] = subTreeXor[first]
			} else {
				// second and first are both subtrees of root
				// 1 segment: root without first and second
				// 2 segment: first
				// 3 segment: second
				segmentXors[0] = subTreeXor[0] ^ subTreeXor[first] ^ subTreeXor[second]
				segmentXors[1] = subTreeXor[first]
				segmentXors[2] = subTreeXor[second]
			}
			// find the score, or difference between largest and smallest segment
			sort.Ints(segmentXors[:])
			diff := segmentXors[2] - segmentXors[0]
			minDiff = min(minDiff, diff)
		}
	}

	return minDiff
}
