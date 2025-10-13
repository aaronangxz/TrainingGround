package main

import "slices"

func removeAnagrams(words []string) []string {
	ans := make([]string, 0, len(words))
	var prevSorted string

	for _, w := range words {
		// If two string equal ignore immediate to avoid sorting
		if w == prevSorted {
			continue
		}

		// Convert to []byte once
		b := []byte(w)
		slices.Sort(b)
		sorted := string(b)

		// Only add if sorted form differs from previous
		if sorted != prevSorted {
			ans = append(ans, w)
			prevSorted = sorted
		}
	}

	return ans
}
