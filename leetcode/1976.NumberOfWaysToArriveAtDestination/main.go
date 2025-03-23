package main

import (
	"container/heap"
	"math"
)

func countPaths(n int, roads [][]int) int {
	const MOD = 1_000_000_007

	graph := make([][][2]int, n)

	for _, road := range roads {
		u, v, time := road[0], road[1], road[2]
		graph[u] = append(graph[u], [2]int{v, time})
		graph[v] = append(graph[v], [2]int{u, time})
	}

	dist := make([]int, n)
	ways := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt64
	}
	dist[0] = 0
	ways[0] = 1

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{node: 0, cost: 0})

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(*Item)

		if cur.cost > dist[cur.node] {
			continue
		}

		for _, neighbor := range graph[cur.node] {
			nextNode, time := neighbor[0], neighbor[1]
			newDist := cur.cost + time

			if newDist < dist[nextNode] {
				dist[nextNode] = newDist
				ways[nextNode] = ways[cur.node]
				heap.Push(pq, &Item{node: nextNode, cost: newDist})
			} else if newDist == dist[nextNode] {
				ways[nextNode] = (ways[nextNode] + ways[cur.node]) % MOD
			}
		}
	}

	return ways[n-1]
}

// Priority Queue Implementation
type Item struct {
	node int
	cost int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].cost < pq[j].cost }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}
