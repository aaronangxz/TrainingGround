package common

import "fmt"

type Node struct {
	Data interface{}
	Next *Node
}

func NewNode(d interface{}) *Node {
	return &Node{
		d,
		nil,
	}
}

func (n *Node) PrintNodes() {
	for n != nil {
		fmt.Print(n.Data)
		n = n.Next
	}
}
