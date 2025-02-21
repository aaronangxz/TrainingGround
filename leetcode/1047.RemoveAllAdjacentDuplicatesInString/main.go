package main

import (
	"fmt"
	"github.com/aaronangxz/TrainingGround/common"
)

func removeDuplicates(s string) string {
	st := common.Stack[rune]{}

	//Push 1st char onto stack first
	st.Push(rune(s[0]))

	//Dispose first char
	s = s[1:]

	for _, ss := range s {
		//Only if stack is not empty, pop when we found another same char
		if !st.IsEmpty() && ss == st.Top() {
			st.Pop()
			continue
		}

		//Otherwise just push onto stack
		st.Push(ss)
	}

	out := ""
	for _, elem := range st.Elements() {
		out += string(elem)
	}
	return out
}

func main() {
	fmt.Println(removeDuplicates("abbaca"))
	fmt.Println(removeDuplicates("azxxzy"))
}
