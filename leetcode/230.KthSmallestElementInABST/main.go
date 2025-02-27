package main

import "github.com/aaronangxz/TrainingGround/common"

func kthSmallest(root *common.TreeNode, k int) int {
	result := -1
	bfs(root, &k, &result)
	return result
}

// Inorder traversal will guaranteed to have values from low to high (properties of BST)
func bfs(root *common.TreeNode, counter, result *int) {
	if root == nil {
		return
	}

	// Run through to the left most node (smallest)
	bfs(root.Left, counter, result)

	// When arriving at the node, decrement counter
	*counter--
	if *counter == 0 {
		// Counter 0, it is now the kth smallest
		*result = root.Val
		return
	}
	// Else check right nodes
	bfs(root.Right, counter, result)
}
