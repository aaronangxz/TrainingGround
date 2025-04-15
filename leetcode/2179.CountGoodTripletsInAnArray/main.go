package main

// Complexity:
// Time O(NLogN) and Space O(N) where N is the length of nums1 and nums2.
func goodTriplets(nums1 []int, nums2 []int) int64 {
	return countIncreasingTriplets(mapPosition(nums1, nums2))
}

// countIncreasingTriplets returns the number of triplets (i, j, k) such that
// i < j < k and nums[i] < nums[j] < nums[k], given that nums is a permutation
// of the integers from 1 to len(nums).
func countIncreasingTriplets(nums []int) int64 {
	// At state j, prefixFwt is the sum Fenwick tree on top of the indicator
	// array, which indicates the existence of numbers for nums[:j].
	// The initial values correspond to the state of j = 0.
	prefixFwt := make(sumFwt, len(nums)+1)

	result := int64(0)
	for j := 1; j < len(nums)-1; j++ {
		// Transition from state j-1 to state j
		prefixFwt.update(nums[j-1], 1)

		// Compute and sum the number of increasing triplets at j
		iChoices := prefixFwt.query(nums[j] - 1)
		kChoices := (len(nums) - nums[j]) - (j - prefixFwt.query(nums[j]))
		result += int64(iChoices) * int64(kChoices)
	}
	return result
}

// mapPosition maps the elements in nums1 to their positions (one-indexed)
// in nums2, given that both nums1 and nums2 are permutations of the integers
// from 0 to len(nums1)-1.
func mapPosition(nums1, nums2 []int) []int {
	valueToIndex2 := make([]int, len(nums2))
	for i, num := range nums2 {
		valueToIndex2[num] = i
	}

	result := make([]int, len(nums1))
	for i, num := range nums1 {
		result[i] = valueToIndex2[num] + 1
	}
	return result
}

// sumFwt is a Fenwick Tree (one-indexed) that supports summing operations.
type sumFwt []int

func (tree *sumFwt) update(index, increment int) {
	for i := index; i < len(*tree); i = (i | (i - 1)) + 1 {
		(*tree)[i] += increment
	}
}

func (tree *sumFwt) query(index int) int {
	result := 0
	for i := index; i > 0; i = i & (i - 1) {
		result += (*tree)[i]
	}
	return result
}
