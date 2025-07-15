package main

func isValid(word string) bool {
	vowels, consonants := 0, 0
	for i := 0; i < len(word); i++ {
		c := word[i]
		if c >= '0' && c <= '9' {
			continue
		}
		if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')) {
			return false
		}
		switch c {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			vowels++
		default:
			consonants++
		}
	}
	return len(word) >= 3 && vowels > 0 && consonants > 0
}
