package main

import "github.com/aaronangxz/TrainingGround/common"

func invertTree(root *common.TreeNode) *common.TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
