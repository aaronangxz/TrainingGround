package main

// Complexity:
// Time O(L*N^2) and Space O(NL) where N and L are the number of words and the average
// length of words, respectively.
func getWordsInLongestSubsequence(words []string, groups []int) []string {
	// dpLen[i] ::= the longest length of the subsequence starting with words[i]
	dpLen := make([]int, len(words))
	// dpNext[i] ::= the seconde index of the subsequence corresponding to dpLen[i]
	dpNext := make([]int, len(words))

	bestStart := len(words) - 1
	dpLen[bestStart] = 1
	dpNext[bestStart] = bestStart

	for start := len(words) - 2; start >= 0; start-- {
		dpLen[start] = 1
		dpNext[start] = start

		for next := start + 1; next+dpLen[start] <= len(words); next++ {
			if groups[next] != groups[start] &&
				dpLen[start] < dpLen[next]+1 &&
				isHammingDistanceOne(words[start], words[next]) {
				dpLen[start] = dpLen[next] + 1
				dpNext[start] = next
			}
		}

		if dpLen[start] > dpLen[bestStart] {
			bestStart = start
		}
	}
	return buildResultSequence(words, bestStart, dpNext)
}

func isHammingDistanceOne(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	diff := 0
	for i := 0; i < len(a) && diff < 2; i++ {
		if a[i] != b[i] {
			diff++
		}
	}
	return diff == 1
}

func buildResultSequence(words []string, bestStart int, dpNext []int) []string {
	result := []string{words[bestStart]}
	curr := bestStart
	for dpNext[curr] != curr {
		curr = dpNext[curr]
		result = append(result, words[curr])
	}
	return result
}
