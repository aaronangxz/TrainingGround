package main

import "github.com/aaronangxz/TrainingGround/common"

func isBalanced(root *common.TreeNode) bool {
	isBalance, _ := dfs(root)
	return isBalance
}

func dfs(root *common.TreeNode) (bool, int) {
	// Root with no child node is considered balance with height 0
	if root == nil {
		return true, 0
	}

	// Recursively solve the following left right tree of current root
	isLeftBalanced, leftHeight := dfs(root.Left)
	isRightBalanced, rightHeight := dfs(root.Right)

	// isLeftBalance -> Left tree is balance
	// isRightBalance -> Right tree is balance
	// abs(leftHeight-rightHeight) <= 1 -> Root is balance
	// Only if all conditions met, root is balance
	isBalanced := isLeftBalanced && isRightBalanced && common.Abs(leftHeight-rightHeight) <= 1

	// Increment 1 to previous max height for current root's height
	return isBalanced, common.Max(leftHeight, rightHeight) + 1
}
