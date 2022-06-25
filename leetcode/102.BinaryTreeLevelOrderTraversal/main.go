package main

import (
	"fmt"
	"github.com/aaronangxz/TrainingGround/common"
)

func levelOrder(root *common.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var out [][]int
	var queue common.Queue

	queue.Enqueue(root)

	for !queue.IsEmpty() {
		size := queue.Size()
		var levelRes []int

		for i := 0; i < size; i++ {
			top := queue.Front().(*common.TreeNode)
			queue.Dequeue()
			levelRes = append(levelRes, top.Val)

			if top.Left != nil {
				queue.Enqueue(top.Left)
			}

			if top.Right != nil {
				queue.Enqueue(top.Right)
			}
		}
		out = append(out, levelRes)
	}
	return out
}

func main() {
	root := common.NewTreeNode(1)
	root.Left = common.NewTreeNode(2)
	root.Right = common.NewTreeNode(3)
	root.Left.Left = common.NewTreeNode(4)
	root.Left.Right = common.NewTreeNode(5)
	root.Right.Left = common.NewTreeNode(6)
	root.Right.Right = common.NewTreeNode(7)
	oL := levelOrder(root)
	fmt.Println(oL)
}
