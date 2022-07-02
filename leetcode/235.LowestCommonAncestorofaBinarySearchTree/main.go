package main

import (
	"github.com/aaronangxz/TrainingGround/common"
)

func lowestCommonAncestorIterative(root, p, q *common.TreeNode) *common.TreeNode {
	for root != nil {
		if p.Val > root.Val && q.Val > root.Val {
			root = root.Left
		} else if p.Val < root.Val && q.Val < root.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return root
}

func lowestCommonAncestorRecursive(root, p, q *common.TreeNode) *common.TreeNode {
	if p.Val < root.Val && q.Val < root.Val {
		//if values of both p and q nodes are less than root, that means it must be in the left subtree
		return lowestCommonAncestorRecursive(root.Left, p, q)
	} else if p.Val > root.Val && q.Val > root.Val {
		//if values of both p and q nodes are greater than root, that means it must be in the right subtree
		return lowestCommonAncestorRecursive(root.Right, p, q)
	}
	//otherwise, the lowest common ancestor is root / p / q
	//since a node can be a descendant of itself according to the LCA definition.
	return root
}

func main() {
	t := common.NewTreeNode(6)
	t.Left = common.NewTreeNode(2)
	t.Right = common.NewTreeNode(8)
	t.Left.Left = common.NewTreeNode(0)
	t.Left.Right = common.NewTreeNode(4)
	t.Left.Right.Left = common.NewTreeNode(3)
	t.Left.Right.Right = common.NewTreeNode(5)
	t.Right.Left = common.NewTreeNode(7)
	t.Right.Right = common.NewTreeNode(9)

	p := t.Left
	q := t.Right
	lowestCommonAncestorRecursive(t, p, q).PrintNode()
	lowestCommonAncestorIterative(t, p, q).PrintNode()
}
