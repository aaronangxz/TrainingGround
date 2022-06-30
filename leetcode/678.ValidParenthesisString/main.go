package main

import (
	"fmt"
	"github.com/aaronangxz/TrainingGround/common"
)

func checkValidString1(s string) bool {
	openP, closeP := 0, 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '*' {
			openP++
		} else {
			openP--
		}

		if s[len(s)-i-1] == ')' || s[len(s)-i-1] == '*' {
			closeP++
		} else {
			closeP--
		}

		if openP < 0 || closeP < 0 {
			return false
		}
	}
	return true
}

func checkValidString(s string) bool {
	minV, maxV := 0, 0

	for _, ss := range s {
		if ss == '(' {
			minV++
			maxV++
		}
		if ss == ')' {
			maxV--
			minV = common.Max(minV-1, 0)
		}
		if ss == '*' {
			maxV++
			minV = common.Max(minV-1, 0)
		}
		if maxV < 0 {
			return false
		}
	}
	return minV == 0
}

func main() {
	s := []string{"()", "(*)", "(*))", "(*)))"}

	for _, ss := range s {
		fmt.Println(checkValidString(ss))
	}
	for _, ss := range s {
		fmt.Println(checkValidString1(ss))
	}
}
