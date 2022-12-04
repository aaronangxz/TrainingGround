package main

import "fmt"

func isVowel(c rune) bool {
	if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
		c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' {
		return true
	}
	return false
}

func halvesAreAlike(s string) bool {
	fH := s[:len(s)/2]
	sH := s[(len(s)-1)/2+1:]

	vF := 0
	vS := 0

	for _, f := range fH {
		if isVowel(f) {
			vF++
		}
	}

	for _, s := range sH {
		if isVowel(s) {
			vS++
		}
	}

	return vF == vS
}

func main() {
	fmt.Println(halvesAreAlike("book"))
	fmt.Println(halvesAreAlike("textbook"))
}
