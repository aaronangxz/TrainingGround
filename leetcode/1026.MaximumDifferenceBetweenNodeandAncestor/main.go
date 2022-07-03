package main

import (
	"fmt"
	"github.com/aaronangxz/TrainingGround/common"
)

func dfs(root *common.TreeNode, mini, maxi int) int {
	if root == nil {
		//return the difference in subtree
		return maxi - mini
	}

	//calculate the current running maximum and minimum
	maxi = common.Max(maxi, root.Val)
	mini = common.Min(mini, root.Val)

	//find the max of both subtree's difference between maxi and mini
	return common.Max(dfs(root.Left, mini, maxi), dfs(root.Right, mini, maxi))
}

func maxAncestorDiff(root *common.TreeNode) int {
	//take current root value as minimum and maximum
	return dfs(root, root.Val, root.Val)
}

func main() {
	n := common.NewTreeNode(8)
	n.Left = common.NewTreeNode(3)
	n.Right = common.NewTreeNode(10)
	n.Left.Left = common.NewTreeNode(1)
	n.Left.Right = common.NewTreeNode(6)
	n.Left.Right.Left = common.NewTreeNode(4)
	n.Left.Right.Right = common.NewTreeNode(7)
	n.Right.Right = common.NewTreeNode(14)
	n.Right.Right.Left = common.NewTreeNode(13)
	fmt.Println(maxAncestorDiff(n))

	n1 := common.NewTreeNode(1)
	n1.Right = common.NewTreeNode(2)
	n1.Right.Right = common.NewTreeNode(0)
	n1.Right.Right.Left = common.NewTreeNode(3)
	fmt.Println(maxAncestorDiff(n1))
}
