package main

import (
	"fmt"
	"github.com/aaronangxz/TrainingGround/common"
)

func makeGood(s string) string {
	st := common.Stack[rune]{}
	st.Push(rune(s[0]))
	s = s[1:]

	for _, c := range s {
		if st.IsEmpty() {
			st.Push(c)
			continue
		}

		if st.Top()-c == 32 || st.Top()-c == -32 {
			st.Pop()
		} else {
			st.Push(c)
		}
	}

	out := ""
	for _, elem := range st.Elements() {
		out += string(elem)
	}
	return out
}

func main() {
	fmt.Println(makeGood("leEeetcode"))
	fmt.Println(makeGood("abBAcC"))
}
