package common

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

func NewTreeNodeWithLeftRight(val int, left, right *TreeNode) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  left,
		Right: right,
	}
}

func NewSliceToTree(values []int) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	root := &TreeNode{Val: values[0]}
	queue := []*TreeNode{root}
	index := 1

	for index < len(values) {
		current := queue[0]
		queue = queue[1:]

		// Add the left child if available
		if index < len(values) {
			current.Left = &TreeNode{Val: values[index]}
			queue = append(queue, current.Left)
			index++
		}

		// Add the right child if available
		if index < len(values) {
			current.Right = &TreeNode{Val: values[index]}
			queue = append(queue, current.Right)
			index++
		}
	}

	return root
}

func NewTreeToSlice(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current != nil {
			result = append(result, current.Val)
			queue = append(queue, current.Left, current.Right)
		} else {
			// Append a placeholder for null nodes (optional; can be omitted if needed)
			result = append(result, -1)
		}
	}

	// Remove trailing -1s (optional, depending on whether you want compact output)
	for len(result) > 0 && result[len(result)-1] == -1 {
		result = result[:len(result)-1]
	}

	return result
}

func (n *TreeNode) PrintNode() {
	if n != nil {
		fmt.Print(n.Val)
	}
	fmt.Println()
}
