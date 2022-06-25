package bfs

import "github.com/aaronangxz/TrainingGround/common"

func OrderLevelTraversal(root *common.TreeNode) [][]int {
	var (
		q   common.Queue
		out [][]int
	)
	if root == nil {
		return [][]int{}
	}

	//root to be in front of the queue
	q.Enqueue(root)

	for !q.IsEmpty() {
		//whatever nodes that exist in the queue now is always on the same level
		//because the next level nodes will only be pushed in later in this loop
		var levelResult []int

		//should take note of length here first, otherwise it might change after enqueue
		size := len(q)

		//loop through all the remaining nodes in the queue (not those new nodes that are being enqueued now)
		for i := 0; i < size; i++ {
			//retrieve the front node
			node := q.Front().(common.TreeNode)
			q.Dequeue()
			//put into level order result
			levelResult = append(levelResult, node.Val)

			if node.Left != nil {
				q.Enqueue(node.Left)
			}

			if node.Right != nil {
				q.Enqueue(node.Right)
			}
		}

		//one level is done, insert into final result
		out = append(out, levelResult)
	}
	return out
}
