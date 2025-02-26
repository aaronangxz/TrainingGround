package main

import "github.com/aaronangxz/TrainingGround/common"

func diameterOfBinaryTree(root *common.TreeNode) int {
	res := 0

	var dfs func(n *common.TreeNode) int
	dfs = func(n *common.TreeNode) int {
		if n == nil {
			return 0
		}

		left := dfs(n.Left)
		right := dfs(n.Right)
		res = max(res, left+right)
		return 1 + max(left, right)
	}
	dfs(root)
	return res
}
