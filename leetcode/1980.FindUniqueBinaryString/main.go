package main

func findDifferentBinaryString(nums []string) string {
	n := len(nums)
	// Make it into a map for easy checking later
	numSet := make(map[string]bool)
	for _, num := range nums {
		numSet[num] = true
	}
	// Begin with an empty string
	return generateBinaryStrings("", numSet, n)
}

func generateBinaryStrings(current string, numSet map[string]bool, n int) string {
	// Base case - when it is long enough + does not exists in map, it is the ans
	// Eventually unwind here when it is len n
	if len(current) == n {
		if !numSet[current] {
			return current
		}
		// No ans
		return ""
	}

	// Try with either one of 0 or 1
	for _, c := range []rune{'0', '1'} {
		// Go multiple levels deep until n
		binaryString := generateBinaryStrings(current+string(c), numSet, n)
		// No ans
		if binaryString != "" {
			return binaryString
		}
	}
	return ""
}
