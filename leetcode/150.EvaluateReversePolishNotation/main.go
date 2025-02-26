package main

import (
	"github.com/aaronangxz/TrainingGround/common"
	"strconv"
)

func evalRPN(tokens []string) int {
	var st common.Stack[int]
	operators := map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}
	for _, token := range tokens {
		if calculate, exist := operators[token]; exist {
			// Get top 2 elements
			b := st.PopTop()
			a := st.PopTop()
			// Push result for later use
			st.Push(calculate(a, b))
		} else {
			// Push new numbers
			num, _ := strconv.Atoi(token)
			st.Push(num)
		}
	}
	return st.Get(0)
}
