package main

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	currentColor := image[sr][sc]
	// Only if the first cell is not the same as the target color, we need to fill the cell
	if currentColor != color {
		fill(image, sr, sc, currentColor, color)
	}
	return image
}

func fill(image [][]int, x int, y int, targetColor, color int) {
	// Check for out of bounds
	if (x < 0 || x > len(image)-1) || (y < 0 || y > len(image[x])-1) {
		return
	}
	// Found matched color as the first cell, update color
	if image[x][y] == targetColor {
		image[x][y] = color
		// Recursively check the adjacent cells
		fill(image, x-1, y, targetColor, color)
		fill(image, x+1, y, targetColor, color)
		fill(image, x, y-1, targetColor, color)
		fill(image, x, y+1, targetColor, color)
	}
}
