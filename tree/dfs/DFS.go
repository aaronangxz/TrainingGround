package dfs

import "github.com/aaronangxz/TrainingGround/common"

func RecursiveDFS(root *common.TreeNode) []*common.TreeNode {
	visited := []*common.TreeNode{}
	return recurse(root, visited)
}

func recurse(root *common.TreeNode, visited []*common.TreeNode) []*common.TreeNode {
	visited = append(visited, root)

	if root.Left != nil {
		visited = recurse(root.Left, visited)
	}

	if root.Right != nil {
		visited = recurse(root.Right, visited)
	}
	return visited
}

func IterativeDfs(root *common.TreeNode) []*common.TreeNode {
	//Stack to store elements to visit
	var s common.Stack

	//Store root node first
	s.Push(root)

	//Keep track of visited nodes
	visited := []*common.TreeNode{}

	//Continue as long as stack is not empty
	for !s.IsEmpty() {
		//Visit the top element in stack
		node := s.Top().(*common.TreeNode)
		visited = append(visited, node)

		//pop the last element
		s.Pop()

		//If right element exists, push onto stack
		if node.Right != nil {
			s.Push(node.Right)
		}

		//If left element exists, push onto stack
		if node.Left != nil {
			s.Push(node.Left)
		}

		//Push right -> left because left node needs to be visited first
	}
	return visited
}
