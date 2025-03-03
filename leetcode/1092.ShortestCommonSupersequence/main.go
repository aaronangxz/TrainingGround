package main

import "strings"

func shortestCommonSupersequence(str1 string, str2 string) string {
	// Step 1: Ensure str1 is the longer string for easier handling
	// If str1 is shorter, swap the roles of str1 and str2
	if len(str1) < len(str2) {
		return shortestCommonSupersequence(str2, str1)
	}

	// Step 2: Create a DP table to store lengths of the longest common subsequences
	// dp[i][j] will represent the length of LCS for str1[i:] and str2[j:]
	dp := make([][]int, len(str1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(str2)+1)
	}

	// Step 3: Fill the DP table starting from the bottom-right corner
	// The idea is to compare characters of str1 and str2 from the end
	// and recursively calculate the LCS length for substrings.
	for i := len(str1) - 1; i >= 0; i-- {
		for j := len(str2) - 1; j >= 0; j-- {
			if str1[i] == str2[j] {
				// If characters match, take the diagonal value and add 1
				dp[i][j] = 1 + dp[i+1][j+1]
			} else {
				// If characters don't match, take the maximum value from either right or down
				dp[i][j] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	// Step 4: Build the result string from the DP table
	// We'll use a strings.Builder to efficiently append characters
	var res strings.Builder
	i, j := 0, 0

	// Step 5: Traverse through str1 and str2 using the DP table to form the supersequence
	// Compare characters from both strings and pick the ones that should go in the result
	for i < len(str1) && j < len(str2) {
		if str1[i] == str2[j] {
			// If characters are the same, add the character to the result and move both pointers
			res.WriteByte(str1[i])
			i++
			j++
		} else if dp[i][j+1] > dp[i+1][j] {
			// If the LCS after skipping str1[i] is greater, pick str2[j] (advance j)
			res.WriteByte(str2[j])
			j++
		} else {
			// Otherwise, pick str1[i] (advance i)
			res.WriteByte(str1[i])
			i++
		}
	}

	// Step 6: Append any remaining characters from both strings
	// After the main loop, one of the strings might still have characters left
	// Append the remaining characters from str1 and str2
	res.WriteString(str1[i:])
	res.WriteString(str2[j:])

	// Return the resulting supersequence as a string
	return res.String()
}
