package common

import (
	"fmt"
)

type BinaryTree struct {
	Node *TreeNode
}

func NewBinaryTree(val int) *BinaryTree {
	return &BinaryTree{Node: NewTreeNode(val)}
}

func (b *BinaryTree) PrintInOrder() {
	var output []int
	inOrderTraversal(b.Node, &output)
	fmt.Println("In-Order:", output)
}

func (b *BinaryTree) PrintPreOrder() {
	var output []int
	preOrderTraversal(b.Node, &output)
	fmt.Println("Pre-Order:", output)
}

func (b *BinaryTree) PrintPostOrder() {
	var output []int
	postOrderTraversal(b.Node, &output)
	fmt.Println("Post-Order:", output)
}

func (b *BinaryTree) PrintLevelOrder() {
	l := orderLevelTraversal(b.Node)
	fmt.Println("Level-Order:", l)
}

func inOrderTraversal(root *TreeNode, output *[]int) {
	if root == nil {
		return
	}
	inOrderTraversal(root.Left, output)
	*output = append(*output, root.Val)
	inOrderTraversal(root.Right, output)
}

func preOrderTraversal(root *TreeNode, output *[]int) {
	if root == nil {
		return
	}
	*output = append(*output, root.Val)
	preOrderTraversal(root.Left, output)
	preOrderTraversal(root.Right, output)
}

func postOrderTraversal(root *TreeNode, output *[]int) {
	if root == nil {
		return
	}
	postOrderTraversal(root.Left, output)
	postOrderTraversal(root.Right, output)
	*output = append(*output, root.Val)
}

func orderLevelTraversal(root *TreeNode) [][]int {
	var (
		q   Queue
		out [][]int
	)
	if root == nil {
		return [][]int{}
	}

	q.Enqueue(root)

	for !q.IsEmpty() {
		var levelResult []int
		size := len(q)

		for i := 0; i < size; i++ {
			node := q.Front().(*TreeNode)
			q.Dequeue()
			levelResult = append(levelResult, node.Val)

			if node.Left != nil {
				q.Enqueue(node.Left)
			}

			if node.Right != nil {
				q.Enqueue(node.Right)
			}
		}
		out = append(out, levelResult)
	}
	return out
}
