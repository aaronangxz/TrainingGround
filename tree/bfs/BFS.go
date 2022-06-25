package bfs

import "github.com/aaronangxz/TrainingGround/common"

func BFS(root *common.TreeNode) []*common.TreeNode {
	var (
		//Maintain a queue of nodes to visit
		q common.Queue
		//Output values of BFS
		out []*common.TreeNode
	)

	//Root will be the first node to visit
	q.Enqueue(root)

	//Continue to visit until queue is empty
	for !q.IsEmpty() {
		//Visit the first node in queue
		node := q.Front().(*common.TreeNode)
		out = append(out, node)

		//If left node of current node is not empty, add into the queue to visit next
		if node.Left != nil {
			q.Enqueue(node.Left)
		}

		//If right node of current node is not empty, add into the queue to visit next
		if node.Right != nil {
			q.Enqueue(node.Right)
		}

		//Pop the first node in the queue (current node)
		q.Dequeue()
	}
	return out
}
