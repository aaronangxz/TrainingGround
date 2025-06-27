package main

import "strings"

/*
what will be done is just brute force, with some optimisation:
- generating all subsequences of length l <= n/k
- all those should have chars not exceeding the max possible freq to achieve the k repetitions
  meaning if total_freq == f; max choosable <= f/k
- just brute force this and try all possibilities, and if equal keep the larger one
- if !isSubsequence don't continue in that recursion branch
*/

func longestSubsequenceRepeatedK(s string, k int) string {
	occ := [26]int{}
	for i := range s {
		occ[int(s[i]-'a')]++
	}
	n := len(s)
	maxStr := ""

	var dfs func(candidate string, currOcc [26]int)
	dfs = func(candidate string, currOcc [26]int) {
		if len(candidate) == n/k {
			return
		}

		for letter := 'z'; letter >= 'a'; letter-- {
			idx := int(letter - 'a')
			if currOcc[idx]+1 > occ[idx]/k {
				continue
			}

			next := candidate + string(letter)
			currOcc[idx]++

			if isSubsequence(s, strings.Repeat(next, k)) {
				maxStr = maxLen(maxStr, next) // prio to maxStr if equality because previous strings are > lexicographically
				dfs(next, currOcc)
			}

			currOcc[idx]--
		}
	}

	dfs("", [26]int{})
	return maxStr
}

func isSubsequence(s, sub string) bool {
	subIdx := 0
	for sIdx := 0; sIdx < len(s) && subIdx < len(sub); sIdx++ {
		if s[sIdx] == sub[subIdx] {
			subIdx++
		}
	}

	return subIdx == len(sub)
}

func maxLen(x, y string) string {
	if len(x) >= len(y) {
		return x
	}
	return y
}
