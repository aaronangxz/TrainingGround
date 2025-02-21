package main

import "github.com/aaronangxz/TrainingGround/common"

type FindElements struct {
	numMap map[int]bool
}

func Constructor(root *common.TreeNode) FindElements {
	num := make(map[int]bool)

	if root != nil {
		root.Val = 0
	}

	var dfs func(root *common.TreeNode)
	dfs = func(root *common.TreeNode) {
		num[root.Val] = true

		if root.Left != nil {
			root.Left.Val = 2*root.Val + 1
			dfs(root.Left)
		}

		if root.Right != nil {
			root.Right.Val = 2*root.Val + 2
			dfs(root.Right)
		}
	}

	dfs(root)

	return FindElements{
		numMap: num,
	}
}

func (this *FindElements) Find(target int) bool {
	return this.numMap[target]
}
