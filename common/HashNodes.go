package common

import "fmt"

type HashNodeList struct {
	Node *HashNodeSingle
	Next *HashNodeList
}

func NewHashNode(key, value int) *HashNodeList {
	return &HashNodeList{
		Node: &HashNodeSingle{
			Key:   key,
			Value: value,
		},
		Next: nil,
	}
}

func (h *HashNodeList) PrintNodes() {
	curr := h

	for {
		if curr == nil {
			fmt.Printf("[]")
			break
		}

		curr.Node.PrintNode()
		if curr.Next != nil {
			fmt.Printf("->")
		}
		curr = curr.Next

		if curr == nil {
			break
		}
	}
	fmt.Println()
}

type HashNodeSingle struct {
	Key   int
	Value int
}

func NewHashNodeSingle(key, value int) *HashNodeSingle {
	return &HashNodeSingle{
		Key:   key,
		Value: value,
	}
}

func (h *HashNodeSingle) PrintNode() {
	curr := h

	if curr == nil {
		fmt.Print("[]")
		return
	}

	if curr.Key == -1 {
		fmt.Print("[Deleted]")
		return
	}

	fmt.Printf("[%v:%v]", curr.Key, curr.Value)
}
