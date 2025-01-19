package main

import (
	"container/heap"
	"math"
)

type Point struct {
	distance float64
	x        int
	y        int
}

type PointHeap []Point

func (h PointHeap) Len() int { return len(h) }
func (h PointHeap) Less(i, j int) bool {
	return h[i].distance > h[j].distance // Max-heap based on distance
}

func (h PointHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PointHeap) Push(x interface{}) {
	*h = append(*h, x.(Point))
}

func (h *PointHeap) Pop() interface{} {
	// Removing and returning the last element while update the existing slice h
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	h := &PointHeap{}
	heap.Init(h)
	for _, point := range points {
		// Simplified form of √(x1 - x2)^2 + (y1 - y2)^2, because (x2, y2) is always 0,0
		distance := math.Sqrt(float64(point[0]*point[0] + point[1]*point[1]))
		heap.Push(h, Point{distance: distance, x: point[0], y: point[1]})

		// Whenever it exceeds k, remove the last element in the heap
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	result := make([][]int, k)
	for i := 0; i < k; i++ {
		tmp := heap.Pop(h).(Point)
		result[i] = []int{tmp.x, tmp.y}
	}
	return result
}
