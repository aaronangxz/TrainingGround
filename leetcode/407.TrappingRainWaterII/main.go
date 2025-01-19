package main

import "container/heap"

func trapRainWater(heightMap [][]int) int {
	visited := make([][]bool, len(heightMap))
	for i := range visited {
		visited[i] = make([]bool, len(heightMap[0]))
	}

	minHeap := &MinHeap{}

	for i := 0; i < len(heightMap); i++ {
		for j := 0; j < len(heightMap[i]); j++ {
			if i == 0 || j == 0 || i == len(heightMap)-1 || j == len(heightMap[i])-1 {
				heap.Push(minHeap, Pos{i, j, heightMap[i][j]})
				visited[i][j] = true
			}
		}
	}

	lastLowestHeight := 0
	ans := 0

	dir := [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	for minHeap.Len() > 0 {
		lowestHeight := heap.Pop(minHeap).(Pos)
		x, y, height := lowestHeight.x, lowestHeight.y, lowestHeight.height

		if height > lastLowestHeight {
			lastLowestHeight = height
		}

		for i := range dir {
			nextX, nextY := x+dir[i][0], y+dir[i][1]
			if nextX < 0 || nextY < 0 || nextX >= len(heightMap) || nextY >= len(heightMap[0]) || visited[nextX][nextY] {
				continue
			}

			visited[nextX][nextY] = true
			if heightMap[nextX][nextY] < lastLowestHeight {
				ans += lastLowestHeight - heightMap[nextX][nextY]
			}
			heap.Push(minHeap, Pos{nextX, nextY, heightMap[nextX][nextY]})
		}
	}

	return ans
}

type Pos struct {
	x      int
	y      int
	height int
}

type MinHeap []Pos

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].height < h[j].height }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Pos))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
