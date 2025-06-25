package main

import (
	"math"
	"sort"
)

func kthSmallestProduct(nums1 []int, nums2 []int, k int64) int64 {
	// Binary search for the k-th smallest product which is eq to the smallest product having countLessOrEqual(mid) >= k.
	lo, hi := int64(-1e10-1), int64(1e10+1)
	ans := hi

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if countLessOrEqual(mid, nums1, nums2) >= k {
			ans = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return ans
}

func countLessOrEqual(val int64, nums1, nums2 []int) int64 {
	var count int64
	for _, n1 := range nums1 {
		if n1 == 0 {
			// If n1 is 0, the product is 0.
			// If val is non-negative, all products with 0 are <= val.
			// If val is negative, no products with 0 are <= val.
			if val >= 0 {
				count += int64(len(nums2))
			}
		} else if n1 > 0 {
			// For positive n1, we need nums2[j] such that nums2[j] <= val / n1.
			// Example: val = -7, n1 = 2. We need nums2[j] <= -3.5. Floor is -4.
			targetFloat := float64(val) / float64(n1)
			targetInt := int(math.Floor(targetFloat))

			// Use sort.Search (binary search) to find the count of elements in nums2 <= targetInt.
			// We're looking for the first element greater than targetInt, so that all before it satisfy the condition
			idx := sort.Search(len(nums2), func(i int) bool {
				return nums2[i] > targetInt
			})
			count += int64(idx)
		} else { // n1 < 0
			// For negative n1, we need nums2[j] such that nums2[j] >= val / n1.
			// Example: val = -7, n1 = -2. We need nums2[j] >= 3.5. Ceil is 4.
			targetFloat := float64(val) / float64(n1)
			targetInt := int(math.Ceil(targetFloat))

			// Use sort.Search (binary search) to find the count of elements in nums2 >= targetInt.
			idx := sort.Search(len(nums2), func(i int) bool {
				return nums2[i] >= targetInt
			})
			count += int64(len(nums2) - idx)
		}
	}
	return count
}
