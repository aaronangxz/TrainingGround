package main

func answerString(word string, n int) string {
	if n == 1 {
		return word
	}

	ans := ""
	m := len(word)
	mxPartLen := m - n + 1

	for i := 0; i < m; i++ {
		tmp := ""
		for j := i; j < min(i+mxPartLen, m); j++ {
			tmp += string(word[j])
		}

		if tmp > ans {
			ans = tmp
		}
	}

	return ans
}
