package common

import "fmt"

type ChildNode struct {
	Data     interface{}
	Children []*ChildNode
}

func NewChildNode(d interface{}) *ChildNode {
	return &ChildNode{
		d,
		nil,
	}
}

func (n *ChildNode) PrintChildNodes() {
	var q Queue
	q.Enqueue(n)
	q.Enqueue(NewChildNode(" Nil "))

	for !q.IsEmpty() {
		top := q.Front().(*ChildNode)
		fmt.Print(top.Data)

		for _, nn := range top.Children {
			q.Enqueue(nn)
		}

		if top.Children != nil {
			q.Enqueue(NewChildNode(" Nil "))
		}

		q.Dequeue()
	}
	fmt.Println()
}
