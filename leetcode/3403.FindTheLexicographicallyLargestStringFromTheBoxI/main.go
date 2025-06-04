package main

func answerString(word string, numFriends int) string {
	if numFriends == 1 {
		return word
	}
	n := len(word) - numFriends + 1
	ans := ""

	for c := 'z'; c >= 'a'; c-- {
		found := false
		for i := 0; i < len(word); i++ {
			if rune(word[i]) != c {
				continue
			}
			found = true
			ans = max(ans, word[i:min(len(word), i+n)])
		}
		if found {
			break
		}
	}

	return ans
}
