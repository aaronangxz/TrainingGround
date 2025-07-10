package main

import (
	"sort"
)

func maxFreeTime(eventTime int, startTime []int, endTime []int) int {
	n := len(startTime)
	// find the 3 biggest gap sizes, start with gap[0], gap[1], gap[n]
	max := []int{startTime[0], startTime[1] - endTime[0],
		eventTime - endTime[n-1]}
	sort.Ints(max)
	// check the rest of the gaps (gap[2] ... gap[n-1]), keep the 3 largest
	for i := 2; i < n; i++ {
		if g := startTime[i] - endTime[i-1]; g > max[0] {
			max[0] = g
			if max[0] > max[1] {
				max[0], max[1] = max[1], max[0]
				if max[1] > max[2] {
					max[1], max[2] = max[2], max[1]
				}
			}
		}
	}

	// calculate the maximum free time by moving the first event
	top := move(startTime[0], startTime[1]-endTime[0],
		endTime[0]-startTime[0], max)
	// calculate the maximum free time by moving the following events
	for i := 1; i < n-1; i++ {
		if m := move(startTime[i]-endTime[i-1], startTime[i+1]-endTime[i],
			endTime[i]-startTime[i], max); m > top {
			top = m
		}
	}
	// calculate the maximum free time by moving the last element
	if m := move(startTime[n-1]-endTime[n-2], eventTime-endTime[n-1],
		endTime[n-1]-startTime[n-1], max); m > top {
		return m
	}
	return top
}

// maximum gap size achievable by moving an event
func move(left, right, event int, max []int) int {
	// what is the biggest gap remaining outside the left and right gaps?
	top := max[2]
	if left == max[2] {
		if right == max[1] {
			top = max[0]
		} else {
			top = max[1]
		}
	} else if right == max[2] {
		if left == max[1] {
			top = max[0]
		} else {
			top = max[1]
		}
	}
	// if we can move the event outside, we can merge it with the gaps
	if top >= event {
		return left + right + event
	}
	// otherwise we can only move the event within left and right
	return left + right
}
