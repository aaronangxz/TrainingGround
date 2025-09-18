package main

import (
	"container/heap"
)

type TaskEntry struct {
	priority int
	taskId   int
}

type MaxHeap []TaskEntry

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	if h[i].priority != h[j].priority {
		return h[i].priority > h[j].priority
	}
	return h[i].taskId > h[j].taskId
}
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(TaskEntry))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type TaskManager struct {
	taskToUser     map[int]int
	taskToPriority map[int]int
	maxHeap        *MaxHeap
}

func Constructor(tasks [][]int) TaskManager {
	tm := TaskManager{
		taskToUser:     make(map[int]int),
		taskToPriority: make(map[int]int),
		maxHeap:        &MaxHeap{},
	}
	heap.Init(tm.maxHeap)

	for _, task := range tasks {
		userId, taskId, priority := task[0], task[1], task[2]
		tm.taskToUser[taskId] = userId
		tm.taskToPriority[taskId] = priority
		heap.Push(tm.maxHeap, TaskEntry{priority, taskId})
	}

	return tm
}

func (this *TaskManager) Add(userId int, taskId int, priority int) {
	this.taskToUser[taskId] = userId
	this.taskToPriority[taskId] = priority
	heap.Push(this.maxHeap, TaskEntry{priority, taskId})
}

func (this *TaskManager) Edit(taskId int, newPriority int) {
	this.taskToPriority[taskId] = newPriority
	heap.Push(this.maxHeap, TaskEntry{newPriority, taskId})
}

func (this *TaskManager) Rmv(taskId int) {
	delete(this.taskToUser, taskId)
	delete(this.taskToPriority, taskId)
}

func (this *TaskManager) ExecTop() int {
	// Lazy cleanup until finding valid task
	for this.maxHeap.Len() > 0 {
		entry := (*this.maxHeap)[0]
		taskId := entry.taskId
		priority := entry.priority

		// Check if task exists and priority is current
		if currentPriority, exists := this.taskToPriority[taskId]; exists && currentPriority == priority {
			// Valid task found, execute it
			heap.Pop(this.maxHeap)
			userId := this.taskToUser[taskId]
			delete(this.taskToUser, taskId)
			delete(this.taskToPriority, taskId)
			return userId
		}

		// Remove outdated entry
		heap.Pop(this.maxHeap)
	}

	return -1 // No valid tasks
}
