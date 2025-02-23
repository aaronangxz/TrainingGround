package main

import "github.com/aaronangxz/TrainingGround/common"

func buildTree(preorder []int, inorder []int) *common.TreeNode {
	// Exhausted elemnts, no new node added
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	// Construct the node at current position,
	// with the root always be at the start of preorder slice
	node := new(common.TreeNode)
	node.Val = preorder[0]

	// Search for the element in the inorder slice, that is the mid point
	// In order is arranged in <left>, <mid>, <right>
	mid := findIdx(preorder[0], inorder)

	// Left subtree
	// Preorder: Exclude idx 0 (that is the root), until mid
	// Inorder: until mid, but exclude mid itself (that is the root)
	node.Left = buildTree(preorder[1:mid+1], inorder[:mid])
	// Right subtree
	// Preorder: From mid onwards to the end
	// Inorder: From mid onwards to the end
	node.Right = buildTree(preorder[mid+1:], inorder[mid+1:])
	return node
}

func findIdx(ele int, arr []int) int {
	for i, a := range arr {
		if ele == a {
			return i
		}
	}
	return -1
}
