package main

func numberOfSubstrings(s string) int {
	l := 0
	res := 0
	count := make(map[string]int)

	// start from l = 0, r = 0
	for r := range len(s) {
		// start counting the char appearing
		count[string(s[r])]++

		// when all of a, b, c were seen
		for len(count) == 3 {
			// Immediately we know there are len(s) - r substrings that is valid
			// i.e. the current substring itself, plus however many that comes up later till the end
			res += len(s) - r

			// Finished checking this left pointer, decrement it
			count[string(s[l])]--

			// If it is 0 after decrementing it, remove from map
			if count[string(s[l])] == 0 {
				delete(count, string(s[l]))
			}
			// Continue to check from the next left pointer
			l++
		}
	}
	return res
}
