package main

import (
	"fmt"
	"github.com/aaronangxz/TrainingGround/common"
)

func isPalindrome(head *common.Node) bool {
	fast, slow := head.Next, head

	//find mid point of list
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	//cut off in the middle
	fast = slow.Next
	slow.Next = nil

	//reverse second half
	var prev, next *common.Node
	curr := fast
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	fast = prev

	//compare between fast and head because we do not have head of slow
	for fast != nil {
		if fast.Data != head.Data {
			return false
		}
		fast = fast.Next
		head = head.Next
	}
	return true
}

func main() {
	n := common.NewNodeFromSlice([]int{1, 2, 2, 1})
	fmt.Println(isPalindrome(n))
	n1 := common.NewNodeFromSlice([]int{1, 2})
	fmt.Println(isPalindrome(n1))
}
