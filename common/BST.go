package common

type BinarySearchTree struct {
	Root *TreeNode
}

func NewBinarySearchTree(val int) *BinarySearchTree {
	return &BinarySearchTree{
		Root: NewTreeNode(val),
	}
}

//WIP
func (b *BinarySearchTree) Insert(val int) *TreeNode {
	if b.Root == nil {
		return NewTreeNode(val)
	}

	if val > b.Root.Val {
		b.Root.Left = b.Insert(val)
	}

	return nil
}
