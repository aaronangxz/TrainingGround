package main

import (
	"fmt"
	"github.com/aaronangxz/TrainingGround/common"
)

func levelOrder(root *common.ChildNode) [][]int {
	//Handle cases when root is nil
	if root == nil {
		return [][]int{}
	}

	var q common.Queue
	//to consolidate results from all levels
	var out [][]int

	//First, enqueue root since this is the first node we will be looking at
	q.Enqueue(root)

	//Continue as long as queue is not empty
	for !q.IsEmpty() {
		//to save results from the current level
		var levelRes []int

		//size of queue is the number of nodes on this level
		//queue can be much larger than this because it consists of nodes from other levels
		size := len(q)

		for i := 0; i < size; i++ {
			//Start looking from the front of the queue
			top := q.Front().(*common.ChildNode)

			//Immediately push the current node value into the current level response
			levelRes = append(levelRes, top.Data.(int))

			//Next, start to look for the nodes on the next level below, i.e. the children
			for i := 0; i < len(top.Children); i++ {
				//For each child, pushes into the queue
				q.Enqueue(top.Children[i])
			}

			//We are done here, dequeue the current node and move on
			q.Dequeue()
		}

		//Lastly, push the result of this level to the overall result slice
		out = append(out, levelRes)
	}
	return out
}

func main() {
	n := common.NewChildNode(1)
	n.Children = append(n.Children, common.NewChildNode(3))
	n.Children = append(n.Children, common.NewChildNode(2))
	n.Children = append(n.Children, common.NewChildNode(4))
	n.Children[0].Children = append(n.Children[0].Children, common.NewChildNode(5))
	n.Children[0].Children = append(n.Children[0].Children, common.NewChildNode(6))
	n.PrintChildNodes()
	fmt.Println(levelOrder(n))
}
