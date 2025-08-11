package main

import (
	"sort"
	"strconv"
)

func convert(n int) string {
	str := strconv.Itoa(n)
	runes := []rune(str)

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	sortedString := string(runes)
	return sortedString
}

func reorderedPowerOf2(n int) bool {
	targetString := convert(n)
	for i := 0; i <= 31; i++ {
		newString := convert(1 << i)
		if newString == targetString {
			return true
		}
	}
	return false
}
