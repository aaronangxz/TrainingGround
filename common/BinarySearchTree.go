package common

type BinarySearchTree struct {
	Tree *BinaryTree
}

func NewBinarySearchTree(val int) *BinarySearchTree {
	return &BinarySearchTree{
		Tree: NewBinaryTree(val),
	}
}

func (b *BinarySearchTree) Insert(val int) {
	insertBSTNode(b.Tree.Node, val)
}

func (b *BinarySearchTree) Delete(val int) {
	deleteBSTNode(b.Tree.Node, val)
}

func insertBSTNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return NewTreeNode(val)
	}
	if val > root.Val {
		root.Right = insertBSTNode(root.Right, val)
	} else if val < root.Val {
		root.Left = insertBSTNode(root.Left, val)
	}
	return root
}

func minNode(root *TreeNode) *TreeNode {
	curr := root

	for curr != nil && curr.Left != nil {
		curr = curr.Left
	}

	return curr
}

//WIP
func deleteBSTNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return root
	}
	if val < root.Val {
		root.Left = deleteBSTNode(root.Left, val)
	} else if val > root.Val {
		root.Right = deleteBSTNode(root.Right, val)
	} else {
		if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Left == nil {

		}
	}
}
