package main

import "github.com/aaronangxz/TrainingGround/common"

func addTwoNumbers(l1 *common.Node, l2 *common.Node) *common.Node {
	//create a new list, summed values will be inserted here
	newList := common.NewNode(0)
	curr := newList
	carry := 0

	for l1 != nil || l2 != nil || carry == 1 {
		sum := 0
		//increment sums from l1 and l2
		if l1 != nil {
			sum += l1.Data.(int)
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Data.(int)
			l2 = l2.Next
		}

		//adding on the carry value
		sum += carry

		//update the next carry value
		carry = sum / 10

		//insert the new sum, modulo by 10
		newNode := common.NewNode(sum % 10)
		curr.Next = newNode
		curr = curr.Next
	}
	return newList.Next
}

func main() {
	addTwoNumbers(common.NewNodeFromSlice([]int{2, 4, 3}), common.NewNodeFromSlice([]int{5, 6, 4})).PrintNodes()
	addTwoNumbers(common.NewNodeFromSlice([]int{0}), common.NewNodeFromSlice([]int{0})).PrintNodes()
	addTwoNumbers(common.NewNodeFromSlice([]int{9, 9, 9, 9, 9, 9, 9}), common.NewNodeFromSlice([]int{9, 9, 9, 9})).PrintNodes()
}
