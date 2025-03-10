package main

func countOfSubstrings(word string, k int) int64 {
	return atLeastKConsonants(word, k) - atLeastKConsonants(word, k+1)
}

func isVowel(b byte) bool {
	return b == 'a' || b == 'o' || b == 'e' || b == 'i' || b == 'u'
}

func atLeastKConsonants(word string, k int) int64 {
	vowel := make(map[byte]int)
	total := int64(0)
	curCons := 0
	l := 0
	for r := 0; r < len(word); r++ {
		if isVowel(word[r]) {
			vowel[word[r]]++
		} else {
			curCons++
		}

		for curCons >= k && len(vowel) >= 5 {
			total += int64(len(word) - r)
			if isVowel(word[l]) {
				vowel[word[l]]--
				if vowel[word[l]] == 0 {
					delete(vowel, word[l])
				}
			} else {
				curCons--
			}
			l++
		}
	}

	return total
}
