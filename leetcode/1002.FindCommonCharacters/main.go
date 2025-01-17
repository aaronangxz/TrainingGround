package main

import "github.com/aaronangxz/TrainingGround/common"

func commonChars(words []string) []string {
	charMap := make([]int, 26)

	// Use the first word as benchmark
	// Safe to do this because the result will at most contain all chars from this word
	for _, c := range words[0] {
		charMap[c-'a']++
	}

	// Starting from the next word
	for i := 1; i < len(words); i++ {
		currCharMap := make([]int, 26)
		// Count the chars of it
		for _, c := range words[i] {
			currCharMap[c-'a']++
		}

		// Compare to the charMap, take the minimum
		// i.e. if there is 0 'a' in this word, then reset 'a' to 0 -> not a common string
		for i := range charMap {
			charMap[i] = common.Min(charMap[i], currCharMap[i])
		}
	}

	var res []string
	// Run through a-z, then append the number of chars accordingly
	for i := 0; i < 26; i++ {
		for j := 0; j < charMap[i]; j++ {
			res = append(res, string(rune(i+'a')))
		}
	}
	return res
}
