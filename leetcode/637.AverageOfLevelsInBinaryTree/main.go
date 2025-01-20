package main

import (
	"fmt"
	"github.com/aaronangxz/TrainingGround/common"
)

func averageOfLevels(root *common.TreeNode) []float64 {
	var q []*common.TreeNode
	var out []float64

	q = append(q, root)

	for len(q) != 0 {
		total := 0
		size := len(q)

		for i := 0; i < size; i++ {
			curr := q[0]
			total += curr.Val
			q = q[1:]

			if curr.Left != nil {
				q = append(q, curr.Left)
			}

			if curr.Right != nil {
				q = append(q, curr.Right)
			}
		}
		avg := float64(total) / float64(size)
		out = append(out, avg)
	}
	return out
}

func averageOfLevelsQueue(root *common.TreeNode) []float64 {
	var q common.Queue
	var out []float64

	q.Enqueue(root)

	for !q.IsEmpty() {
		total := 0
		size := q.Size()

		for i := 0; i < size; i++ {
			curr := q.Front().(*common.TreeNode)
			total += curr.Val
			q.Dequeue()

			if curr.Left != nil {
				q.Enqueue(curr.Left)
			}

			if curr.Right != nil {
				q.Enqueue(curr.Right)
			}
		}
		avg := float64(total) / float64(size)
		out = append(out, avg)
	}
	return out
}

func main() {
	root := common.NewTreeNode(3)
	root.Left = common.NewTreeNode(9)
	root.Right = common.NewTreeNode(20)
	root.Right.Left = common.NewTreeNode(15)
	root.Right.Right = common.NewTreeNode(7)
	aL := averageOfLevels(root)
	aLQ := averageOfLevelsQueue(root)
	fmt.Println(aL, aLQ)
}
