package main

import (
	"github.com/aaronangxz/TrainingGround/common"
	"strings"
)

func wordSubsets(words1 []string, words2 []string) []string {
	// Map to keep track of char frequency in words2
	charMap := make(map[string]int)
	var ans []string

	for _, w2 := range words2 {
		for _, c := range w2 {
			// Do this because if there are similar words in words2
			// e.g. ["oo", "ooo"] should be taken as 3 "o" (max), not 5 "o" (sum)
			charMap[string(c)] = common.Max(charMap[string(c)], strings.Count(w2, string(c)))
		}
	}

	for _, w1 := range words1 {
		allMatch := true
		charCount := make(map[string]int)
		// Do the same for words1
		for _, c := range w1 {
			charCount[string(c)] = common.Max(charCount[string(c)], strings.Count(w1, string(c)))
		}

		for c, count := range charMap {
			// As long as it is lesser than that in words2, it is not a match
			if charCount[c] < count {
				allMatch = false
				break
			}
		}
		if allMatch {
			ans = append(ans, w1)
		}
	}
	return ans
}
