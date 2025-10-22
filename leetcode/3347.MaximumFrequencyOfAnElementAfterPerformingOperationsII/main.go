package main

import "sort"

func maxFrequency(nums []int, k int, numOperations int) int {
	sort.Ints(nums)
	freq := map[int]int{}
	for _, v := range nums {
		freq[v]++
	}

	// for a certain candidate, this returns how many values can match it with numOperations, while taking in consideration if the candidate already exists with a certain freq in the array
	resultForCandidate := func(candidate int) int {
		first, last := binarySearch(nums, k, candidate, true), binarySearch(nums, k, candidate, false)
		totalItemsToChange := last - first + 1 - freq[candidate]
		return freq[candidate] + min(totalItemsToChange, numOperations)
	}

	res, lo := 0, 0
	for hi := 0; hi < len(nums); hi++ {
		// handle sliding window where you assume we do one operation for each item in the window, and we find the max window where: nums[hi] - nums[lo] <= 2 * k; meaning all those numbers can be a single value
		for lo <= hi && (nums[hi]-nums[lo] > 2*k || hi-lo+1 > numOperations) {
			lo++
		}
		res = max(res, hi-lo+1, resultForCandidate(nums[hi]))
	}

	return res
}

func binarySearch(nums []int, k, val int, shouldFindFirst bool) int {
	res := len(nums)
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if abs(nums[mid]-val) <= k {
			res = mid
			if shouldFindFirst {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else {
			if nums[mid]-val > k {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		}
	}

	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
