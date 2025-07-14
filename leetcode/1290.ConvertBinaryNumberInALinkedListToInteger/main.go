package main

import "github.com/aaronangxz/TrainingGround/common"

func getDecimalValue(head *common.TreeNodeWithNext) int {
	result := 0
	for head != nil {
		result = result*2 + head.Val
		head = head.Next
	}
	return result
}
