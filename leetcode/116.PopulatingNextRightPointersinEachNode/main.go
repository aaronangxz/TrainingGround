package main

import (
	"fmt"
	"github.com/aaronangxz/TrainingGround/common"
)

func connect(root *common.TreeNodeWithNext) *common.TreeNodeWithNext {
	if root == nil {
		return root
	}

	var q common.Queue
	q.Enqueue(root)

	for !q.IsEmpty() {
		var connectNode *common.TreeNodeWithNext
		size := len(q)
		for i := 0; i < size; i++ {
			top := q.Front().(*common.TreeNodeWithNext)
			q.Dequeue()

			//connect node to its next
			top.Next = connectNode
			//update right connecting node for next node
			connectNode = top

			//push nodes in next level
			if top.Right != nil {
				q.Enqueue(top.Right)
			}
			if top.Left != nil {
				q.Enqueue(top.Left)
			}
		}
	}
	return root
}

func main() {
	n := common.NewTreeNodeWithNext(1)
	n.Left = common.NewTreeNodeWithNext(2)
	n.Right = common.NewTreeNodeWithNext(3)
	n.Left.Left = common.NewTreeNodeWithNext(4)
	n.Left.Right = common.NewTreeNodeWithNext(5)
	n.Right.Left = common.NewTreeNodeWithNext(6)
	n.Right.Right = common.NewTreeNodeWithNext(7)

	fmt.Println(connect(n))
}
