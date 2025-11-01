package main

import "github.com/aaronangxz/TrainingGround/common"

func modifiedList(nums []int, head *common.TreeNodeWithNext) *common.TreeNodeWithNext {
	toDelete := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		toDelete[v] = struct{}{}
	}

	dummy := &common.TreeNodeWithNext{Val: 0, Next: head}
	prev := dummy
	curr := head

	for curr != nil {
		if _, ok := toDelete[curr.Val]; ok {
			prev.Next = curr.Next
		} else {
			prev = curr
		}
		curr = curr.Next
	}
	return dummy.Next
}
