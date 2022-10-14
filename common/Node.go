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
	var nodes []interface{}

	for n != nil {
		nodes = append(nodes, n.Data)
		n = n.Next
	}

	for i, node := range nodes {
		fmt.Print(node)
		if i != len(nodes)-1 {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}

func NewNodeFromSlice(d []int) *Node {
	if len(d) == 0 {
		return nil
	}

	head := NewNode(d[0])

	if len(d) == 1 {
		return head
	}

	curr := head
	for i := 1; i < len(d); i++ {
		n := NewNode(d[i])
		curr.Next = n
		curr = curr.Next
	}
	return head
}
