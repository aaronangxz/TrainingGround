package main

import (
	"fmt"
	"github.com/aaronangxz/TrainingGround/common"
)

func smallestNumber(pattern string) string {
	var res string
	st := common.Stack[string]{}

	for i := 0; i < len(pattern)+1; i++ {
		st.Push(fmt.Sprintf("%v", i+1))

		for !st.IsEmpty() && (i == len(pattern) || string(pattern[i]) == "I") {
			res += st.Top()
			st.Pop()
		}
	}
	return res
}
