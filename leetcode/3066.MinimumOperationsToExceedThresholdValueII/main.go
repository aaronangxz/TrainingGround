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

func minOperations(nums []int, k int) int {
	mHeap := &minHeap{}
	heap.Init(mHeap)
	for _, n := range nums {
		if n < k {
			heap.Push(mHeap, n)
		}
	}
	ops := 0
	for mHeap.Len() > 0 {
		first := heap.Pop(mHeap).(int)
		if mHeap.Len() == 0 {
			ops++
			break
		}
		second := heap.Pop(mHeap).(int)
		now := min(first, second)*2 + max(first, second)
		if now < k {
			heap.Push(mHeap, now)
		}
		ops++
	}
	return ops
}

func checkK(mHeap *minHeap, k int) bool {
	for _, m := range *mHeap {
		if m < k {
			return false
		}
	}
	return true
}
