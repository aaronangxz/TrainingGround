package main

import "container/heap"

type Task struct {
	userId   int
	taskId   int
	priority int
}

type TaskHeap []Task

func (h TaskHeap) Len() int { return len(h) }
func (h TaskHeap) Less(i, j int) bool {
	if h[i].priority == h[j].priority {
		return h[i].taskId > h[j].taskId // higher taskId wins
	}
	return h[i].priority > h[j].priority // higher priority wins
}
func (h TaskHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *TaskHeap) Push(x any)   { *h = append(*h, x.(Task)) }
func (h *TaskHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type TaskManager struct {
	h      TaskHeap
	record map[int]Task
}

func Constructor(tasks [][]int) TaskManager {
	tm := TaskManager{
		h:      TaskHeap{},
		record: make(map[int]Task),
	}
	heap.Init(&tm.h)
	for _, t := range tasks {
		tm.Add(t[0], t[1], t[2])
	}
	return tm
}

func (tm *TaskManager) Add(userId, taskId, priority int) {
	task := Task{userId, taskId, priority}
	heap.Push(&tm.h, task)
	tm.record[taskId] = task
}

func (tm *TaskManager) Edit(taskId, newPriority int) {
	old := tm.record[taskId]
	updated := Task{old.userId, taskId, newPriority}
	heap.Push(&tm.h, updated)
	tm.record[taskId] = updated
}

func (tm *TaskManager) Rmv(taskId int) {
	delete(tm.record, taskId)
}

func (tm *TaskManager) ExecTop() int {
	for tm.h.Len() > 0 {
		top := heap.Pop(&tm.h).(Task)
		latest, ok := tm.record[top.taskId]
		if !ok {
			continue
		}
		if latest.priority != top.priority {
			continue
		}
		delete(tm.record, top.taskId)
		return top.userId
	}
	return -1
}
