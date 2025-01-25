package main

import "sort"

func lexicographicallySmallestArray(nums []int, limit int) []int {
	numsSorted := make([]int, len(nums))
	// Sort the nums slice first, make a copy so the original slice is not affected
	copy(numsSorted, nums)
	sort.Ints(numsSorted)

	// Keep track of the current group
	currGroup := 0
	numToGroup := make(map[int]int)
	// Using this as a slice of queues
	groupToList := make(map[int][]int)

	// Pre-fill for the first element
	numToGroup[numsSorted[0]] = currGroup
	groupToList[currGroup] = append(groupToList[currGroup], numsSorted[0])

	// Starting from the next
	for i := 1; i < len(nums); i++ {
		// If it is beyond the limit, make a new group
		if abs(numsSorted[i]-numsSorted[i-1]) > limit {
			currGroup++
		}

		// Update group mapping of this element
		numToGroup[numsSorted[i]] = currGroup
		// Add to the group
		groupToList[currGroup] = append(groupToList[currGroup], numsSorted[i])
	}

	for i := 0; i < len(nums); i++ {
		num := nums[i]

		// From the original slice, get the group via mapping
		group := numToGroup[num]
		// Update the slice by popping from the queue front
		nums[i] = groupToList[group][0]
		// De-queue operation
		groupToList[group] = groupToList[group][1:]
	}

	return nums
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
