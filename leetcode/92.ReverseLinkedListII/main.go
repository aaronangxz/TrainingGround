package main

import "github.com/aaronangxz/TrainingGround/common"

func reverseBetween(head *common.Node, left int, right int) *common.Node {
	curr := head
	var prev, next *common.Node
	prev.Next = head
	i := 0

	for i < right {
		//move pointer to the left position first
		if i < left {
			prev = curr
			curr = curr.Next
			i++
		} else {
			next = curr.Next
			curr.Next = curr.Next.Next
			next.Next = prev.Next
			prev.Next = next
		}
	}
	return head
}

func main() {
	reverseBetween(common.NewNodeFromSlice([]int{1, 2, 3, 4, 5}), 2, 4).PrintNodes()
	//reverseBetween(common.NewNodeFromSlice([]int{5}), 1, 1).PrintNodes()
}
