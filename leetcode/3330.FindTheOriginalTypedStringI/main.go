package main

func possibleStringCount(word string) int {
	ans := 1
	cnt := 1
	ch := word[0]

	for i := 1; i < len(word); i++ {
		if word[i] == ch {
			cnt++
		} else {
			ans += cnt - 1
			ch = word[i]
			cnt = 1
		}
	}

	ans += cnt - 1
	return ans
}
