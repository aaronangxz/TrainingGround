package main

import (
	"fmt"
	"github.com/aaronangxz/TrainingGround/common"
)

func dfs(root *common.TreeNode, k int, m map[int]bool) bool {
	if root == nil {
		return true
	}

	if m[k-root.Val] {
		return true
	}
	m[root.Val] = true

	return dfs(root.Left, k, m) || dfs(root.Right, k, m)
}

func findTarget(root *common.TreeNode, k int) bool {
	return dfs(root, k, map[int]bool{})
}

func main() {
	t := common.NewTreeNode(5)
	t.Left = common.NewTreeNode(3)
	t.Right = common.NewTreeNode(6)
	t.Left.Left = common.NewTreeNode(2)
	t.Left.Right = common.NewTreeNode(4)
	t.Right.Right = common.NewTreeNode(7)
	fmt.Println(findTarget(t, 9))
}
