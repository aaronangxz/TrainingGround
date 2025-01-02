package main

func vowelStrings(words []string, queries [][]int) []int {
	var resp []int

	// Offset by 1, because we need an empty space in index 0,
	// in case of calculating for 1st element
	prefixSum := make([]int, len(words)+1)

	// Keeping track of the previous sum occurred
	prevSum := 0

	// Build prefix sum slice
	for i, w := range words {
		// Whenever we see a vowel string, increment by 1
		if isVowel(w) {
			prevSum++
		}
		// Remember we offset by 1
		prefixSum[i+1] = prevSum
	}

	// Actual query check
	for _, q := range queries {
		l, r := q[0], q[1]
		// Look for the sum of (end + 1) - start
		// This will cover the sum of the selected range
		resp = append(resp, prefixSum[r+1]-prefixSum[l])
	}

	return resp
}

func isVowel(src string) bool {
	start := string(src[0])
	end := string(src[len(src)-1])
	return isCharVowel(start) && isCharVowel(end)
}

func isCharVowel(src string) bool {
	return src == "a" || src == "e" || src == "i" || src == "o" || src == "u"
}
