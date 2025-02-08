package main

import "container/heap"

type minHeap []int

func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

type NumberContainers struct {
	numToIndices map[int]*minHeap
	indexToNum   map[int]int
}

func Constructor() NumberContainers {
	return NumberContainers{
		numToIndices: make(map[int]*minHeap),
		indexToNum:   make(map[int]int),
	}
}

// Time O(logn)
func (this *NumberContainers) Change(index int, number int) {
	this.indexToNum[index] = number
	if this.numToIndices[number] == nil {
		this.numToIndices[number] = new(minHeap)
	}
	heap.Push(this.numToIndices[number], index)
}

// Time O(klogn)
func (this *NumberContainers) Find(number int) int {
	if this.numToIndices[number] == nil {
		return -1
	}

	for this.numToIndices[number].Len() != 0 {
		current := (*this.numToIndices[number])[0]
		if this.indexToNum[current] == number {
			return current
		}
		_ = heap.Pop(this.numToIndices[number]).(int)
	}
	return -1
}

/**
 * Your NumberContainers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Change(index,number);
 * param_2 := obj.Find(number);
 */
