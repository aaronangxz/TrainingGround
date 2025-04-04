package main

import "github.com/aaronangxz/TrainingGround/common"

func lcaDeepestLeaves(root *common.TreeNode) *common.TreeNode {
	_, res := dfs(root)

	return res
}

func dfs(root *common.TreeNode) (int, *common.TreeNode) {
	left := 0
	lNode := &common.TreeNode{}
	right := 0
	rNode := &common.TreeNode{}

	if root.Left != nil {
		left, lNode = dfs(root.Left)
	}
	if root.Right != nil {
		right, rNode = dfs(root.Right)
	}

	if left == right {
		return left + 1, root
	}
	if left > right {
		return left + 1, lNode
	}
	return right + 1, rNode
}
