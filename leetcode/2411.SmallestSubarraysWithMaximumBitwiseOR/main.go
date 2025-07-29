package main

/*
res[i] is the length of the minimum sized subarray starting at i with maximum bitwise OR.
--
optimal way:( (O(n) time, O(32) == O(1) space)
- loop right to left, while tracking leftMostIdx of each bit from i..len-1
- each iteration we update the indices
- then to guarantee a maxOr, we just check what is the furthest index from i out of the first times a bit was set

example:
 0 1 2 3 4
[1,0,2,1,3]
at 2,1,3 we will have leftMostIdx after 2 for bit0:3; for bit1:2
thus the endIdx=3; and we don't need the fourth element (3) as bits were already set
*/

func smallestSubarrays(nums []int) []int {
	leftMostIdx, res := [32]int{}, make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		endIdx := i
		for bitIdx := 0; bitIdx < 32; bitIdx++ {
			mask := 1 << bitIdx
			if nums[i]&mask > 0 {
				leftMostIdx[bitIdx] = i
			}
			endIdx = max(endIdx, leftMostIdx[bitIdx])
		}
		res[i] = endIdx - i + 1
	}
	return res
}
