package main

import (
	"github.com/aaronangxz/TrainingGround/common"
)

func swapNodes(head *common.Node, k int) *common.Node {
	fast, slow, curr := head, head, head

	//slow pointer will go k steps first, as well as curr pointer
	for k > 1 {
		slow = slow.Next
		curr = curr.Next
		k--
	}

	//at this point slow is at kth position
	//continue to move forward until the end of list, at the same time move the fast pointer
	//it will reach the last kth position
	for curr.Next != nil {
		curr = curr.Next
		fast = fast.Next
	}

	//swap values
	fast.Data, slow.Data = slow.Data, fast.Data
	return head
}

func main() {
	swapNodes(common.NewNodeFromSlice([]int{1, 2, 3, 4, 5}), 2).PrintNodes()
	swapNodes(common.NewNodeFromSlice([]int{7, 9, 6, 6, 7, 8, 3, 0, 9, 5}), 5).PrintNodes()
}
