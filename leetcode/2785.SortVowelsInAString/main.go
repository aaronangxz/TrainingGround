package main

import "slices"

func isVowel(ch byte) bool {
	return ch == 'a' || ch == 'u' || ch == 'o' || ch == 'i' || ch == 'e' ||
		ch == 'A' || ch == 'U' || ch == 'O' || ch == 'I' || ch == 'E'
}

func sortVowels(s string) string {
	chars := []byte{}
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
			chars = append(chars, s[i])
		}
	}

	slices.Sort(chars)
	idx := 0
	bytes := []byte(s)

	for i := 0; i < len(bytes); i++ {
		if isVowel(bytes[i]) {
			bytes[i] = chars[idx]
			idx++
		}
	}

	return string(bytes)
}
