package main

func findAnagrams(s string, p string) []int {
	var res []int
	patternS := make([]int, 26)
	slidingS := make([]int, 26)

	if len(s) < len(p) {
		return res
	}

	// Collect all char of p, and s that is as long as p
	for i := 0; i < len(p); i++ {
		patternS[p[i]-'a']++
		slidingS[s[i]-'a']++
	}

	// If it is already equal here, means the first p is an anagram,
	// hence index 0 is an answer
	// [x,x,x,d,e,f,g,h]
	if isSliceEqual(patternS, slidingS) {
		res = append(res, 0)
	}

	// Starting from the index outside of p
	for i := len(p); i < len(s); i++ {
		// Start sliding R to the R + 1
		slidingS[s[i]-'a']++
		// Drop the char outside of L
		slidingS[s[i-len(p)]-'a']--
		// [a,x,x,x,e,f,g,h] -> drop a, add d -> [b,c,d]
		// [a,b,x,x,x,f,g,h] -> drop b, add e -> [c,d,e]
		// [a,b,c,x,x,x,g,h] -> drop c, add f -> [d,e,f]
		// ...

		// If the rolling window still has the same chars, they are anagrams
		// Add the starting index to the answer (the next idx from where we dropped)
		if isSliceEqual(patternS, slidingS) {
			res = append(res, i-len(p)+1)
		}
	}
	return res
}

func isSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, aa := range a {
		if b[i] != aa {
			return false
		}
	}
	return true
}
