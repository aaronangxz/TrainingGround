package main

import (
	"fmt"
	"github.com/aaronangxz/TrainingGround/common"
	"math"
)

func maxGain(root *common.TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}

	//eliminate negatives
	//in the case where left or right child is negative
	//parent node can choose not to go there. hence, we treat it as 0
	l := common.Max(maxGain(root.Left, maxSum), 0)
	r := common.Max(maxGain(root.Right, maxSum), 0)

	//max path sum with SPLIT
	//self + the total of left and right subtree, where left and right doesn't split
	//because in recursive returns, values are always non-split
	//i.e. only the current node is splitting now
	*maxSum = common.Max(*maxSum, root.Val+l+r)
	fmt.Println(*maxSum)

	//max path sum without SPLIT
	//self + the largest path in either left or right
	return root.Val + common.Max(l, r)
}

func maxPathSum(root *common.TreeNode) int {
	maxSum := math.MinInt32
	maxGain(root, &maxSum)
	return maxSum
}

func main() {
	n := common.NewTreeNode(-10)
	n.Left = common.NewTreeNode(9)
	n.Right = common.NewTreeNode(20)
	n.Right.Left = common.NewTreeNode(15)
	n.Right.Right = common.NewTreeNode(7)

	fmt.Println(maxPathSum(n))
}
