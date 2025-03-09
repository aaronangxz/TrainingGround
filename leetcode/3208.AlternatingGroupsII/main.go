package main

func numberOfAlternatingGroups(colors []int, groupSize int) int {
	left := 0        // Left pointer to track the start of the current window
	right := 0       // Right pointer to expand the window
	result := 0      // Count of valid alternating groups
	n := len(colors) // Length of the colors array

	// Iterate through the array to find valid alternating groups
	for left < n {
		// Move the right pointer forward
		right++

		// Check if the current tile and the previous tile have the same color
		if colors[right%n] == colors[(right-1)%n] {
			// If they are the same, the alternating sequence breaks
			// Move the left pointer to the current position to start a new window
			left = right
		} else if right-left+1 == groupSize {
			// If the window size equals the required group size, it's a valid group
			result++ // Increment the result count
			left++   // Move the left pointer to slide the window forward
		}
	}

	return result
}
