package main

// Complexity:
// Time O(t+len(s)) and Space O(1).
func lengthAfterTransformations(s string, t int) int {
	const MODULE = 1_000_000_007

	freq := make([]int, 26)
	for _, c := range s {
		freq[c-'a']++
	}

	head := 0 // head points to the index that corresponds to a
	for range t {
		newHead := head - 1
		if newHead < 0 {
			newHead += 26
		}

		freq[head] = (freq[head] + freq[newHead]) % MODULE
		head = newHead
	}

	result := 0
	for _, f := range freq {
		result = (result + f) % MODULE
	}
	return result
}
