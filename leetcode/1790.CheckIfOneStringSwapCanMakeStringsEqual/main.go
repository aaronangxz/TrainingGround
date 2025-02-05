package main

func areAlmostEqual(s1 string, s2 string) bool {
	// Keep track of the idx where diff is observed
	firstIdx, secondIdx := 0, 0
	// And the total of diffs
	unequalCount := 0
	for i := range len(s1) {
		if s1[i] != s2[i] {
			// Increment the diff count
			unequalCount++
			// There should not be more than 2 diffs
			if unequalCount > 2 {
				return false
			}
			// Found the first diff, record the idx
			if unequalCount == 1 {
				firstIdx = i
			}
			// Found the second diff, record the idx
			if unequalCount == 2 {
				secondIdx = i
			}
		}
	}
	// If the string is identical if swapped either way, it is equal
	return s1[firstIdx] == s2[secondIdx] && s1[secondIdx] == s2[firstIdx]
}
