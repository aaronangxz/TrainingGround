package main

import (
	"fmt"
	"github.com/aaronangxz/TrainingGround/common"
)

func recurse(root *common.ChildNode, out *[]int) {
	if root == nil {
		return
	}

	//append root node first
	*out = append(*out, root.Data.(int))

	//and for each child in node, recurse individually
	for _, c := range root.Children {
		recurse(c, out)
	}
}

func preorderRecursive(root *common.ChildNode) []int {
	var out []int
	recurse(root, &out)
	return out
}

func preorderIterative(root *common.ChildNode) []int {
	st := common.Stack[*common.ChildNode]{}
	var out []int

	if root == nil {
		return out
	}

	//push root node onto the stack first
	st.Push(root)

	//continue whenever stack is not empty
	for !st.IsEmpty() {
		//retrieve the top element, save into result and pop it
		top := st.Top()
		out = append(out, top.Data.(int))
		st.Pop()

		//for each child of the top element, push onto the stack
		//start from the back because FIFO, we want to print from the left onwards
		for i := len(top.Children) - 1; i >= 0; i-- {
			st.Push(top.Children[i])
		}
	}
	return out
}

func main() {
	n := common.NewChildNode(1)
	n.Children = append(n.Children, common.NewChildNode(3))
	n.Children = append(n.Children, common.NewChildNode(2))
	n.Children = append(n.Children, common.NewChildNode(4))
	n.Children[0].Children = append(n.Children[0].Children, common.NewChildNode(5))
	n.Children[0].Children = append(n.Children[0].Children, common.NewChildNode(6))
	n.PrintChildNodes()
	fmt.Println(preorderRecursive(n))
	fmt.Println(preorderIterative(n))
}
