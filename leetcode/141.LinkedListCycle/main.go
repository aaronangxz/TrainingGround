package main

import "github.com/aaronangxz/TrainingGround/common"

func hasCycle(head *common.TreeNodeWithNext) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}
