package main

func canBeValid(s string, locked string) bool {
	n := len(s)
	// If the length of the string is odd, it cannot be a valid string
	// There will be no full match of parentheses
	if n%2 != 0 {
		return false
	}

	openCount := 0
	for i := 0; i < n; i++ {
		// From left to right, treat the open and unlocked parentheses as open parentheses
		if s[i] == '(' || locked[i] == '0' {
			openCount++
		} else {
			openCount--
		}
		// If the number of open parentheses is less than 0, it means that there are more close parentheses than open parentheses, which is not valid
		if openCount < 0 {
			return false
		}
	}

	closeCount := 0
	for i := n - 1; i >= 0; i-- {
		// From right to left, treat the close and unlocked parentheses as close parentheses
		if s[i] == ')' || locked[i] == '0' {
			closeCount++
		} else {
			closeCount--
		}
		// If the number of close parentheses is less than 0, it means that there are more open parentheses than close parentheses, which is not valid
		if closeCount < 0 {
			return false
		}
	}

	return true
}
