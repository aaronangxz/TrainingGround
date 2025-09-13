package main

import "strings"

func maxFreqSum(s string) int {
	freq := make(map[rune]int)
	con, vow := 0, 0
	for _, c := range s {
		freq[c]++
	}
	for ch, count := range freq {
		if strings.ContainsRune("aeiou", ch) {
			if count > vow {
				vow = count
			}
		} else {
			if count > con {
				con = count
			}
		}
	}
	return con + vow
}
