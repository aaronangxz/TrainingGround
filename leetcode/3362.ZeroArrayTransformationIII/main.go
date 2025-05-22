package main

import "container/heap"

func maxRemoval(nums []int, queries [][]int) int {
	queryByStartIdx := make([][][]int, len(nums))
	for _, q := range queries {
		queryByStartIdx[q[0]] = append(queryByStartIdx[q[0]], q)
	}
	availableQueries := NewHeap() // returns the one with the highest endIdx out of the available queries
	endEvents := make([]int, len(nums)+1)
	currCapacity := 0
	for i := range nums {
		// add queries available for selection to the heap
		for _, q := range queryByStartIdx[i] {
			heap.Push(availableQueries, q)
		}
		// remove selected queries that ended
		currCapacity -= endEvents[i]
		// keep picking the best queries (the ones ending last) until matching the current element
		for availableQueries.Len() > 0 && currCapacity < nums[i] {
			poppedQuery := heap.Pop(availableQueries).([]int)
			endEvents[poppedQuery[1]+1]++
			if poppedQuery[1] >= i { // only add capacity if the available query adds value at this point
				currCapacity++
			}
		}
		// if we picked and we couldnt make it return -1
		if currCapacity < nums[i] {
			return -1
		}
	}

	return availableQueries.Len()
}

//--------heap implementation--------------------

// An IntHeap is a min-heap of ints.
type IntHeap [][]int

func NewHeap() *IntHeap {
	return &IntHeap{}
}

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i][1] > h[j][1] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([]int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
