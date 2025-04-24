package main

func countCompleteSubarrays(nums []int) int {
	dis := make(map[int]bool)
	for _, v := range nums {
		dis[v] = true
	}
	k := len(dis)

	freq := make(map[int]int)
	l, res := 0, 0
	for r := 0; r < len(nums); r++ {
		freq[nums[r]]++
		for len(freq) == k {
			res += len(nums) - r
			freq[nums[l]]--
			if freq[nums[l]] == 0 {
				delete(freq, nums[l])
			}
			l++
		}
	}
	return res
}
