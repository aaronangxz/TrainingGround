package common

type Node struct {
	data interface{}
	next *Node
}

func NewNode(d interface{}) *Node {
	return &Node{
		d,
		nil,
	}
}
