package main

func longestPalindrome(words []string) int {
	cnt := make(map[string]int, len(words))
	for _, w := range words {
		cnt[w]++
	}
	res := 0
	for w, c := range cnt {
		rev := string([]byte{w[1], w[0]})
		if w < rev {
			if oc, ok := cnt[rev]; ok {
				m := c
				if oc < m {
					m = oc
				}
				res += 4 * m
			}
		}
	}
	for w, c := range cnt {
		if w[0] == w[1] {
			res += 4 * (c / 2)
		}
	}
	for w, c := range cnt {
		if w[0] == w[1] && c%2 == 1 {
			res += 2
			break
		}
	}
	return res
}
