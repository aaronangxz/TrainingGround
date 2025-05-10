package main

func minSum(nums1 []int, nums2 []int) int64 {
	var n01, n02, sum1, sum2 int64
	for _, n := range nums1 {
		sum1 += int64(n)
		if n == 0 {
			n01++
		}
	}
	for _, n := range nums2 {
		sum2 += int64(n)
		if n == 0 {
			n02++
		}
	}
	if (n01 == 0 && sum1 < sum2+n02) || (n02 == 0 && sum2 < sum1+n01) {
		return -1
	}
	return max(sum1+n01, int64(sum2+n02))
}
