package main

import "github.com/aaronangxz/TrainingGround/common"

func revInOrder(root *common.TreeNode, pre *int) *common.TreeNode {
	//inorder binary_tree, but start from the right subtree
	if root.Right != nil {
		revInOrder(root.Right, pre)
	}

	//increment the current node value + previous value
	root.Val = root.Val + *pre
	//update the previous value
	*pre = root.Val

	if root.Left != nil {
		revInOrder(root.Left, pre)
	}
	return root
}

func bstToGst(root *common.TreeNode) *common.TreeNode {
	pre := 0
	return revInOrder(root, &pre)
}

func main() {
	n := common.NewTreeNode(4)
	n.Left = common.NewTreeNode(1)
	n.Right = common.NewTreeNode(6)
	n.Left.Left = common.NewTreeNode(0)
	n.Left.Right = common.NewTreeNode(2)
	n.Left.Right.Right = common.NewTreeNode(3)
	n.Right.Left = common.NewTreeNode(5)
	n.Right.Right = common.NewTreeNode(7)
	n.Right.Right.Right = common.NewTreeNode(8)
	bstToGst(n).PrintNode()
}
