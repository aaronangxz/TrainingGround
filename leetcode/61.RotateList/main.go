package main

import "github.com/aaronangxz/TrainingGround/common"

func rotateRight(head *common.Node, k int) *common.Node {
	if head == nil {
		return head
	}

	n := head
	length := 1

	//make it into a circular linked list
	for n.Next != nil {
		n = n.Next
		length++
	}
	n.Next = head

	//save moves with modulo
	k %= length

	//if k % length is 0, we are at the old head hence don't need to move
	if k > 0 {
		//move pointer to the node before the new rotated head
		for i := 0; i < length-k; i++ {
			n = n.Next
		}
	}

	//point to the new head
	newHead := n.Next
	//disconnect the circular linked list
	n.Next = nil

	return newHead
}

func main() {
	n := common.NewNodeFromSlice([]int{1, 2, 3, 4, 5})
	rotateRight(n, 2).PrintNodes()
}
