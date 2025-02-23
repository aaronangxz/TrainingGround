package main

import "github.com/aaronangxz/TrainingGround/common"

func constructFromPrePost(preorder []int, postorder []int) *common.TreeNode {
	idxMap := make(map[int]int)
	for i, p := range postorder {
		idxMap[p] = i
	}

	var build func(i1, i2, j int) *common.TreeNode
	build = func(i1, i2, j int) *common.TreeNode {
		if i1 > i2 {
			return nil
		}

		root := new(common.TreeNode)
		root.Val = preorder[i1]
		if i1 != i2 {
			leftVal := preorder[i1+1]
			mid := idxMap[leftVal]
			leftSize := mid - j + 1
			root.Left = build(i1+1, i1+leftSize, j)
			root.Right = build(i1+leftSize+1, i2, mid+1)
		}
		return root
	}
	return build(0, len(preorder)-1, 0)
}
