package main

import (
	"fmt"
	"github.com/aaronangxz/TrainingGround/common"
)

func makeGood(s string) string {
	var st common.Stack
	st.Push(rune(s[0]))
	s = s[1:]

	for _, c := range s {
		if st.IsEmpty() {
			st.Push(c)
			continue
		}

		if st.Top().(rune)-c == 32 || st.Top().(rune)-c == -32 {
			st.Pop()
		} else {
			st.Push(c)
		}
	}

	out := ""
	for _, sts := range st {
		out += string(sts.(rune))
	}
	return out
}

func main() {
	fmt.Println(makeGood("leEeetcode"))
	fmt.Println(makeGood("abBAcC"))
}
