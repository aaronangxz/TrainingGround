package main

func countGood(nums []int, k int) int64 {
	var t int64 = 0
	var total int64 = 0
	seen := make(map[int]int)
	n := int64(len(nums))
	j := 0
	for i := 0; i < len(nums); i++ {
		seen[nums[i]]++
		t += int64(seen[nums[i]] - 1)
		for j < i && t > int64(k-1) {
			t -= int64(seen[nums[j]] - 1)
			seen[nums[j]]--
			j++
		}
		total += int64(i - j + 1)
	}
	return n*(n+1)/2 - total
}
