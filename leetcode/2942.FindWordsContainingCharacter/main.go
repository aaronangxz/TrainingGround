package main

import "strings"

func findWordsContaining(words []string, x byte) []int {
	res := make([]int, 0, len(words))
	for i, w := range words {
		if strings.IndexByte(w, x) >= 0 {
			res = append(res, i)
		}
	}
	return res
}
