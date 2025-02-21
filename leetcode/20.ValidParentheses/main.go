package main

import (
	"fmt"
	"github.com/aaronangxz/TrainingGround/common"
)

func isValidParentheses(s string) bool {
	st := common.Stack[rune]{}

	for _, ss := range s {
		//push into stack if it is an opening parenthesis
		if ss == '(' || ss == '[' || ss == '{' {
			st.Push(ss)
		} else {
			//cannot proceed if there aren't any open parenthesis
			if st.IsEmpty() {
				return false
			}

			//get the top element
			top := st.Top()

			//if any of these doesn't match
			if (ss == ')' && top != '(') || (ss == ']' && top != '[') || (ss == '}' && top != '{') {
				return false
			}

			//means it matches, pop it off
			st.Pop()
		}
	}
	return st.IsEmpty()
}

func main() {
	s := "()[]{}"
	fmt.Println(isValidParentheses(s))

	s1 := "([{}])[]"
	fmt.Println(isValidParentheses(s1))

	s2 := "[){}()"
	fmt.Println(isValidParentheses(s2))
}
