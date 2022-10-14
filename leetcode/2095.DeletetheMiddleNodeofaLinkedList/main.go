package main

import "github.com/aaronangxz/TrainingGround/common"

func deleteMiddle(head *common.Node) *common.Node {
	if head.Next == nil {
		return nil
	}

	//When there are only 2 nodes, keep the first
	if head.Next.Next == nil {
		head.Next = nil
		return head
	}

	slow, fast, prev := head, head, head

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	//slow will reach the middle node
	prev.Next = slow.Next

	return head
}

func main() {
	l := common.NewNodeFromSlice([]int{1, 3, 4, 7, 1, 2, 6})
	deleteMiddle(l).PrintNodes()

	l1 := common.NewNodeFromSlice([]int{1, 2, 3, 4})
	deleteMiddle(l1).PrintNodes()

	l2 := common.NewNodeFromSlice([]int{2, 1})
	deleteMiddle(l2).PrintNodes()
}
