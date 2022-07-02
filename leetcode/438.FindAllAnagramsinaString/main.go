package main

import "fmt"

func isSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, aa := range a {
		if b[i] != aa {
			return false
		}
	}
	return true
}

func findAnagrams(s string, p string) []int {
	var res []int
	patternMap, slidingMap := make([]int, 26), make([]int, 26)

	if len(s) < len(p) {
		return res
	}

	//fill both maps with p
	for i := 0; i < len(p); i++ {
		patternMap[p[i]-'a']++
		slidingMap[s[i]-'a']++
	}

	//if both are identical, that means there is one anagram at the 0 index
	if isSliceEqual(patternMap, slidingMap) {
		res = append(res, 0)
	}

	//start from the next index after p
	for i := len(p); i < len(s); i++ {
		//slide window to the right
		slidingMap[s[i]-'a']++

		//retract one index from the left
		slidingMap[s[i-len(p)]-'a']--

		if isSliceEqual(patternMap, slidingMap) {
			//since the first index is now i-len(p)+1, after the window was slided
			res = append(res, i-len(p)+1)
		}
	}
	return res
}

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
	fmt.Println(findAnagrams("abab", "ab"))
}
