package main

import "sort"

func minimumDeletions(word string, k int) int {
	arr := make([]int, 26)
	del := 0
	ans := len(word)
	for _, ch := range word {
		arr[ch-'a']++
	}
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		res := del
		minFreq := arr[i]
		for j := len(arr) - 1; j > i; j-- {
			if arr[j]-minFreq <= k {
				break
			}
			res += arr[j] - minFreq - k
		}
		ans = min(ans, res)
		del += arr[i]
	}
	return ans
}
