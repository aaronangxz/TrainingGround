package main

func kLengthApart(nums []int, k int) bool {
	lastOccured := -1
	for i, num := range nums {
		if num == 1 {
			if lastOccured != -1 {
				gap := i - lastOccured - 1
				if gap < k {
					return false
				}
			}
			lastOccured = i
		}
	}
	return true
}
