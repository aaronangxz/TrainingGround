package main

import (
	"fmt"
	"github.com/aaronangxz/TrainingGround/common"
)

func main() {
	t := common.NewTreeNode(1)
	t.Left = common.NewTreeNode(2)
	t.Right = common.NewTreeNode(3)
	t.Left.Left = common.NewTreeNode(4)
	fmt.Println(tree2str(t))

	t1 := common.NewTreeNode(1)
	t1.Left = common.NewTreeNode(2)
	t1.Right = common.NewTreeNode(3)
	t1.Left.Right = common.NewTreeNode(4)
	fmt.Println(tree2str(t1))

	t2 := common.NewTreeNode(1)
	t2.Left = common.NewTreeNode(2)
	t2.Right = common.NewTreeNode(3)
	t2.Left.Left = common.NewTreeNode(4)
	t2.Left.Right = common.NewTreeNode(5)
	t2.Right.Left = common.NewTreeNode(6)
	t2.Right.Right = common.NewTreeNode(7)
	fmt.Println(tree2str(t2))
}

func tree2str(root *common.TreeNode) string {
	if root == nil {
		return ""
	}

	left, right := "", ""

	if root.Left != nil {
		left = "(" + tree2str(root.Left) + ")"
	}

	if root.Right != nil {
		left = "(" + tree2str(root.Left) + ")"
		right = "(" + tree2str(root.Right) + ")"
	}

	return fmt.Sprintf("%v", root.Val) + left + right
}
