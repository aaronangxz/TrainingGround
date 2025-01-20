package main

import "github.com/aaronangxz/TrainingGround/common"

func removeNthFromEnd(head *common.Node, n int) *common.Node {
	fast, slow := head, head

	//fast node will move to the nth position first
	for n > 0 {
		fast = fast.Next
		n--
	}

	//if fast has reached nil, we need to remove the head
	if fast == nil {
		return head.Next
	}

	//move fast to the end, while moving slow as well
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	//slow is at the previous node of the node to be removed
	//disconnect and link to the next next node
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
