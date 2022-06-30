package main

import (
	"fmt"
	"github.com/aaronangxz/TrainingGround/common"
)

func isValidBST(root *common.TreeNode) bool {
	return dfs(root, nil, nil)
}

func dfs(root, min, max *common.TreeNode) bool {
	//it's fine when there's no more nodes
	if root == nil {
		return true
	}

	//when respective subtrees violate properties of BST
	if (min != nil && root.Val <= min.Val) || (max != nil && root.Val >= max.Val) {
		return false
	}

	//keeping track of min / max for left / right trees
	//for left tree, subtrees should be < root, hence max is root
	//for right tree, subtrees should be > root, hence min is root
	return dfs(root.Left, min, root) && dfs(root.Right, root, max)
}

func isValidBSTInOrder(root *common.TreeNode) bool {
	var n *common.TreeNode
	return inOrder(root, &n)
}

func inOrder(root *common.TreeNode, prev **common.TreeNode) bool {
	if root == nil {
		return true
	}

	if !inOrder(root.Left, prev) {
		return false
	}

	if *prev != nil && (*prev).Val >= root.Val {
		return false
	}
	*prev = root
	return inOrder(root.Right, prev)
}

func main() {
	n := common.NewTreeNode(5)
	n.Left = common.NewTreeNode(1)
	n.Right = common.NewTreeNode(4)
	n.Right.Left = common.NewTreeNode(3)
	n.Right.Right = common.NewTreeNode(6)
	fmt.Println(isValidBST(n))
	fmt.Println(isValidBSTInOrder(n))

	n1 := common.NewTreeNode(2)
	n1.Left = common.NewTreeNode(1)
	n1.Right = common.NewTreeNode(3)
	fmt.Println(isValidBST(n1))
	fmt.Println(isValidBSTInOrder(n1))

	n2 := common.NewTreeNode(1)
	n2.Left = common.NewTreeNode(1)
	fmt.Println(isValidBST(n2))
	fmt.Println(isValidBSTInOrder(n2))
}
