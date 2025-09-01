package main

import "container/heap"

type Class struct {
	pass, total int
	gain        float64
}

type MaxHeap []Class

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].gain > h[j].gain }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(Class)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func gain(pass, total int) float64 {
	return float64(pass+1)/float64(total+1) - float64(pass)/float64(total)
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	h := &MaxHeap{}
	for _, c := range classes {
		heap.Push(h, Class{c[0], c[1], gain(c[0], c[1])})
	}
	heap.Init(h)

	for extraStudents > 0 {
		top := heap.Pop(h).(Class)
		top.pass++
		top.total++
		top.gain = gain(top.pass, top.total)
		heap.Push(h, top)
		extraStudents--
	}

	sum := 0.0
	for _, c := range *h {
		sum += float64(c.pass) / float64(c.total)
	}
	return sum / float64(len(classes))
}
