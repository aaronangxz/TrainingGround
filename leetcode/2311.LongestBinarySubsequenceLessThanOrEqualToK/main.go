package main

func longestSubsequence(s string, k int) int {
	res := 0
	cur := 0
	for i := len(s) - 1; i >= 0; i-- {
		c := s[i]
		if c == '0' {
			res++
		} else if res < 31 && cur+(1<<res) <= k {
			cur += 1 << res
			res++
		}
	}
	return res
}
