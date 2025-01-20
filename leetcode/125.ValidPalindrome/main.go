package main

import "fmt"

func toLower(b byte) int {
	c := int(b)
	if c >= 'A' && c <= 'Z' {
		c += 32
	}
	return c
}

func isMatch(l, r byte) bool {
	//before comparing, convert all to lowercase
	c1, c2 := toLower(l), toLower(r)
	return c1 == c2
}

func isAlphaNumeric(b byte) bool {
	c := int(b)
	return (c >= '0' && c <= '9') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= 'a' && c <= 'z')
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		//fast-forward and skip all non-alphanumeric characters
		for l < r && !isAlphaNumeric(s[l]) {
			l++
		}
		for l < r && !isAlphaNumeric(s[r]) {
			r--
		}
		if !isMatch(s[l], s[r]) {
			return false
		}
		l++
		r--
	}
	return true
}

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}
