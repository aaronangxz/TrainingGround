package main

import "sort"

func minimumBoxes(apple []int, capacity []int) int {
	sum := 0
	for _, ap := range apple {
		sum += ap
	}
	sort.Ints(capacity)
	i := len(capacity) - 1
	for sum > 0 {
		sum -= capacity[i]
		i--
	}
	return len(capacity) - 1 - i
}
