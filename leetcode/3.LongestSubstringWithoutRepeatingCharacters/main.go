package main

func lengthOfLongestSubstring(s string) int {
	lastSeenCharAtIdx := make(map[byte]int)
	longest := 0
	start := 0

	for end := 0; end < len(s); end++ {
		curr := s[end]
		// If seen before and it is later than the previous sliding window start
		if idx, ok := lastSeenCharAtIdx[curr]; ok && idx > start {
			// use it as the new sliding window start (jump there, window is then smaller now)
			start = idx
		}
		// Update the longest string thus far
		longest = max(longest, end-start+1)
		// Record last seen idx with + 1 so later we can access the next in line element directly
		lastSeenCharAtIdx[curr] = end + 1
	}
	return longest
}
