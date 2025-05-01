package main

import "sort"

// BrunoT - user8184Aj - https://leetcode.com/problems/maximum-number-of-tasks-you-can-assign/solutions/6704952/nice-binary-search-problem-for-resource-allocation
func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
	n, m := len(tasks), len(workers)
	queue := make([]int, min(n, m)) // reuse for check
	sort.Ints(tasks)
	sort.Ints(workers)

	left, right := 1, min(n, m)
	for left <= right {
		mid := (left + right) / 2
		if check(tasks, workers, pills, strength, mid, queue) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}

func check(tasks, workers []int, pills, strength, mid int, queue []int) bool {
	p := pills
	m := len(workers)
	first, last := 0, 0
	ptr := m - 1

	for i := mid - 1; i >= 0; i-- {
		// Fill queue with workers that can solve this task (with or without pill)
		for ptr >= m-mid && workers[ptr]+strength >= tasks[i] {
			queue[last] = workers[ptr]
			last++
			ptr--
		}

		if first == last {
			return false // no available workers
		} else if queue[first] >= tasks[i] {
			// strongest worker can do the task
			first++
			continue
		}

		if p == 0 {
			return false // no pills left
		}
		p--
		last-- // use pill on weakest that can help
	}
	return true
}
