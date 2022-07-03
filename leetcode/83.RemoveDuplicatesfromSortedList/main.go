package main

import "github.com/aaronangxz/TrainingGround/common"

func deleteDuplicates(head *common.Node) *common.Node {
	if head == nil {
		return head
	}

	//pointer to insert nodes
	curr := head
	//keep track of the previous node value
	prev := head.Data
	//pointer to iterate through
	n := head.Next

	for n != nil {
		//unique node
		if n.Data != prev {
			//insert into next position
			curr.Next = n
			//update previous value
			prev = n.Data
			//move head pointer
			curr = curr.Next
		}
		//always move to the next node
		n = n.Next
	}

	//no more nodes, the remaining should be cut off
	curr.Next = nil
	return head
}

func main() {
	n := common.NewNodeFromSlice([]int{1, 1, 2})
	deleteDuplicates(n).PrintNodes()
	n1 := common.NewNodeFromSlice([]int{1, 1, 2, 3, 3})
	deleteDuplicates(n1).PrintNodes()
}
