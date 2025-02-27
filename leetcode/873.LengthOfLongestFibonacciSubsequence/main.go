package main

func lenLongestFibSubseq(arr []int) int {
	hmap := make(map[int]struct{})
	for _, v := range arr {
		hmap[v] = struct{}{}
	}

	var dp func(first, second, count int) int
	dp = func(first, second, count int) int {
		if _, ok := hmap[first+second]; ok {
			if count == 0 {
				count = 3
			} else {
				count++
			}
			return dp(second, first+second, count)
		}
		return count
	}

	ans := 0
	for i := range len(arr) {
		for j := i + 1; j < len(arr); j++ {
			ans = max(ans, dp(arr[i], arr[j], 0))
		}
	}

	return ans
}
