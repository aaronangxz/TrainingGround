package main

func minOperations(nums []int, k int) int {
	st := make(map[int]struct{})
	for _, x := range nums {
		if x < k {
			return -1
		} else if x > k {
			st[x] = struct{}{}
		}
	}
	return len(st)
}
