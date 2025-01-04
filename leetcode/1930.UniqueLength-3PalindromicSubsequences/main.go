package main

func countPalindromicSubsequence(s string) int {
	// res: to store the unique results in the form of first 2 char of string
	// since we can define the palindrome string just from it
	// e.g. "ab" we know that it is -> "aba"
	res := map[string]struct{}{}
	// left: keep track of the char on the left of 'mid', not count sensitive
	left := map[string]struct{}{}
	// right: keep track of the char on the right of 'mid', count sensitive
	right := make(map[string]int)

	// Since we start on idx 0, everything is on the right
	// Fill right with all char counts
	for _, c := range s {
		right[string(c)]++
	}

	for _, c := range s {
		// c is viewed as mid
		cStr := string(c)
		// Decrement c from the right, since it is now used as mid
		if _, ok := right[cStr]; ok {
			right[cStr]--
		}

		// Match whatever in the left of mid to check if it is also in right
		// 1x char from left forms first char of palindrome
		// cStr forms the middle char of palindrome
		// 1x char from right forms last char of palindrome
		for l := range left {
			// There is a match + it is not exhausted yet
			if count, found := right[l]; found && count > 0 {
				// Record the first 2 char of palindrome string
				res[l+cStr] = struct{}{}
			}
		}
		// Once everything is done, current mid is now converted to left for next idx
		left[cStr] = struct{}{}
	}
	// Return unique results
	return len(res)
}
