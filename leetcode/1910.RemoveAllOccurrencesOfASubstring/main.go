package main

func removeOccurrences(s string, part string) string {
	// Iterate through the index (difference of s and part)
	for i := 0; i <= len(s)-len(part); i++ {
		// If the first n index of s is equals to part
		if s[i:i+len(part)] == part {
			// Remove the substring
			s = s[0:i] + s[i+len(part):]
			// Reverse i back to 0
			i = -1
		}
	}
	return s
}
