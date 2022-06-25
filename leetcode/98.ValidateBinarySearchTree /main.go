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

func main() {
	n := common.NewTreeNode(5)
	n.Left = common.NewTreeNode(1)
	n.Right = common.NewTreeNode(4)
	n.Right.Left = common.NewTreeNode(3)
	n.Right.Right = common.NewTreeNode(6)
	fmt.Println(isValidBST(n))

	n1 := common.NewTreeNode(2)
	n1.Left = common.NewTreeNode(1)
	n1.Right = common.NewTreeNode(3)
	fmt.Println(isValidBST(n1))
}
