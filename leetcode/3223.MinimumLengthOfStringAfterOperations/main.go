package main

func minimumLength(s string) int {
	m := make(map[string]int)

	// Count all the chars
	for _, c := range s {
		m[string(c)]++

		// Whenever a given char has 3 occurrences, it means that we can delete 1 in front and 1 at the back
		if m[string(c)] == 3 {
			m[string(c)] -= 2
		}
	}

	count := 0
	// Count the remaining chars
	for _, v := range m {
		count += v
	}
	return count
}
