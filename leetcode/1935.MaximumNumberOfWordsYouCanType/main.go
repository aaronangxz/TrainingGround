package main

import "strings"

func canBeTypedWords(text string, brokenLetters string) int {
	words := strings.Split(text, " ")
	count := 0
	for _, word := range words {
		canType := true
		for _, c := range brokenLetters {
			if strings.ContainsRune(word, c) {
				canType = false
				break
			}
		}
		if canType {
			count++
		}
	}
	return count
}
