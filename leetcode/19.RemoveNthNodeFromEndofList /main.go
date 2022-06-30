package main

import "github.com/aaronangxz/TrainingGround/common"

func removeNthFromEnd(head *common.Node, n int) *common.Node {
	fast, slow := head, head

	for n > 0 {
		fast = fast.Next
		n--
	}

	if fast == nil {
		return head.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}

func main() {
	n := common.NewNode(1)
	n.Next = common.NewNode(2)
	n.Next.Next = common.NewNode(3)
	n.Next.Next.Next = common.NewNode(4)
	n.Next.Next.Next.Next = common.NewNode(5)

	res := removeNthFromEnd(n, 2)
	res.PrintNodes()
}
