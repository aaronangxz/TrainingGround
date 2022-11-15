package main

import (
	"fmt"
	"github.com/aaronangxz/TrainingGround/common"
)

func countNodes(root *common.TreeNode) int {
	cnt := 0
	count(root, &cnt)
	return cnt
}

// Simple In order to count nodes
func count(root *common.TreeNode, cnt *int) {
	if root == nil {
		return
	}

	if root.Left != nil {
		count(root.Left, cnt)
	}

	*cnt++

	if root.Right != nil {
		count(root.Right, cnt)
	}
}

func main() {
	t := common.NewTreeNode(1)
	t.Left = common.NewTreeNode(2)
	t.Left.Left = common.NewTreeNode(4)
	t.Left.Right = common.NewTreeNode(5)
	t.Right = common.NewTreeNode(3)
	t.Right.Left = common.NewTreeNode(6)
	fmt.Println(countNodes(t))

	t1 := common.NewTreeNode(1)
	fmt.Println(countNodes(t1))
}
