package main

import (
	"fmt"
	"github.com/aaronangxz/TrainingGround/common"
)

func validateStackSequencesConstSpace(pushed []int, popped []int) bool {
	i, j := 0, 0

	for _, p := range pushed {
		fmt.Println("before", pushed)
		pushed[i+1] = p
		fmt.Println("aft", pushed)
		for i > 0 && (pushed[i-1] == popped[j]) {
			i--
			j++
		}
	}
	return i == 0
}

func validateStackSequencesWithStack(pushed []int, popped []int) bool {
	var st common.Stack
	i := 0

	//simulate push and pop of a stack and compare with the popped slice
	for _, p := range pushed {
		st.Push(p)

		//when the top is equals to the element in popped slice
		//pop the stack and continue to check for the next element
		for !st.IsEmpty() && st.Top().(int) == popped[i] {
			st.Pop()
			i++
		}
	}

	//sequence is correct is stack is empty in the end
	return st.IsEmpty()
}

func validateStackSequencesWithSlice(pushed []int, popped []int) bool {
	var st []int
	i := 0

	for _, p := range pushed {
		st = append(st, p)
		for len(st) != 0 && st[len(st)-1] == popped[i] {
			st = st[:len(st)-1]
			i++
		}
	}
	return len(st) == 0
}

func main() {
	fmt.Println(validateStackSequencesWithSlice([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
	fmt.Println(validateStackSequencesWithSlice([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}))
	fmt.Println(validateStackSequencesConstSpace([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
}
