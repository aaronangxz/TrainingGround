package common

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
