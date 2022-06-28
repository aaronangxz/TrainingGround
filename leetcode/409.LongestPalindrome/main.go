package main

import "fmt"

func longestPalindrome(s string) int {
	m := make(map[int]int)
	longest := 0
	hasOdd := false

	//fill in all the characters and occurrences
	for i := range s {
		m[int(s[i])]++
	}

	for _, val := range m {
		//for each character, only sum up half of it to eliminate those with single occurrence
		longest += val / 2

		//if any of them has a single element, we can use it as the centre char
		if val%2 == 0 {
			hasOdd = true
		}
	}

	//double the length so we get the full palindrome string
	longest *= 2

	//add in the centre char if applicable
	if hasOdd {
		longest++
	}
	return longest
}

func main() {
	s := "abccccdd"
	fmt.Println(longestPalindrome(s))
}
