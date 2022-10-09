package main

import (
	"github.com/aaronangxz/TrainingGround/common"
)

func addNode(root *common.TreeNode, val int, depth int, now int) *common.TreeNode {
	if root == nil {
		return nil
	}

	//Found correct depth
	//Connect to left / right respectively
	if now == depth {
		root.Left = common.NewTreeNodeWithLeftRight(val, root.Left, nil)
		root.Right = common.NewTreeNodeWithLeftRight(val, nil, root.Right)
		return root
	}

	//solve left / right subtree respectively
	root.Left = addNode(root.Left, val, depth, now+1)
	root.Right = addNode(root.Right, val, depth, now+1)
	return root
}

func addOneRow(root *common.TreeNode, val int, depth int) *common.TreeNode {
	//Case when depth is 1, simply create new root and connect whole tree to left subtree
	if depth == 1 {
		return common.NewTreeNodeWithLeftRight(val, root, nil)
	}
	//start with now = 1
	//depth-1 because we need to insert node when we arrive at one level above of the depth
	return addNode(root, val, depth-1, 1)
}

func main() {
	t := common.NewTreeNode(4)
	t.Left = common.NewTreeNode(2)
	t.Right = common.NewTreeNode(6)
	t.Left.Left = common.NewTreeNode(3)
	t.Left.Right = common.NewTreeNode(1)
	t.Right.Left = common.NewTreeNode(5)
	ans := addOneRow(t, 1, 2)
	b := common.BinaryTree{Node: ans}
	b.PrintLevelOrder()
}
