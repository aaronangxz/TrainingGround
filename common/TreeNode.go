package common

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

func (n *TreeNode) PrintNode() {
	if n != nil {
		fmt.Print(n.Val)
	}
	fmt.Println()
}
