package main

func queryResults(limit int, queries [][]int) []int {
	// Maps a ball to its current color.
	ballToColor := make(map[int]int)
	// Maps a color to the count of balls painted with that color.
	colorCounts := make(map[int]int)
	// Number of distinct colors currently in use.
	distinctColors := 0
	// Prepare the result slice.
	result := make([]int, 0, len(queries))

	for _, query := range queries {
		ball, newColor := query[0], query[1]

		// If the ball was previously painted, decrement the count for the old color.
		if oldColor, exists := ballToColor[ball]; exists {
			colorCounts[oldColor]--
			if colorCounts[oldColor] == 0 {
				distinctColors--
			}
		}

		// Paint the ball with the new color and increment the corresponding count.
		ballToColor[ball] = newColor
		colorCounts[newColor]++
		if colorCounts[newColor] == 1 {
			distinctColors++
		}

		result = append(result, distinctColors)
	}
	return result
}
