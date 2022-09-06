package main

/**
* Strong password checker
* A password is strong if:
* has 5 to 20 character
* Contains at least 1 lowercase letter, 1 uppercase letter and 1 digit
* Does not contains >=3 repeating letter such as ..aaa. or bbb.... or ....cccc but ..aa.a is ok
*
* Given a password, find the minimum action needed to make the given password strong.
* In 1 action you can either:
* Delete a character from the password
* Insert a character in any position in the password
*
*Input: a, Output: 4 // as you can insert 4 character, 1 Upper case letter , 1 digit and another character
*Input: aA1, Output: 3 // Same reason as above
*Input: 11Abc, Output: 0 // It already a strong password
 */

import "fmt"

//WIP
func main() {
	fmt.Println(passwordChecker("aA1"))
}

func isContainsCase(s string) (bool, int) {
	hasUpper, hasLower, hasDigit := false, false, false
	for _, char := range s {
		if char >= 65 && char <= 90 {
			hasUpper = true
		} else if char >= 97 && char <= 122 {
			hasLower = true
		} else if char >= 48 && char <= 57 {
			hasDigit = true
		}
	}

	count := 0
	if hasUpper {
		count++
	}

	if hasLower {
		count++
	}

	if hasDigit {
		count++
	}
	return hasUpper && hasLower && hasDigit, count
}

func isRepeated(s string) (bool, int) {
	repeated := false
	repeatedPos := 0
	count := 1
	prev := s[0]

	for i := 1; i < len(s); i++ {
		if s[i] == prev {
			count++
			if count == 3 {
				repeatedPos++
				repeated = true
			}
		} else {
			count = 1
		}
		prev = s[i]
	}
	return repeated, repeatedPos
}

func passwordChecker(s string) int {
	count := 0

	if len(s) <= 5 {
		return 5 - len(s)
	} else if len(s) >= 20 {
		count += 20 - len(s)
	}

	if ok, n := isContainsCase(s); !ok {
		count += 3 - n
	}

	if repeated, n := isRepeated(s); repeated {
		count += n
	}
	return count
}
