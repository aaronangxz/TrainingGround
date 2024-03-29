package main

import (
	"fmt"
	"github.com/aaronangxz/TrainingGround/common"
	"github.com/aaronangxz/TrainingGround/tree/bfs"
	"github.com/aaronangxz/TrainingGround/tree/dfs"
)

func main() {
	root := common.NewBinarySearchTree(1)
	root.Insert(2)
	root.Insert(3)
	root.Insert(4)
	root.Insert(5)
	root.Tree.PrintInOrder()
}

func testBST() {

}

func testTraversal() {
	root := common.NewBinaryTree(1)
	root.Node.Left = common.NewTreeNode(2)
	root.Node.Right = common.NewTreeNode(3)
	root.Node.Left.Left = common.NewTreeNode(4)
	root.Node.Left.Right = common.NewTreeNode(5)

	root.PrintPreOrder()
	root.PrintInOrder()
	root.PrintPostOrder()
}

func testBFSDFS() {
	//     1
	//   2   3
	//  4 5 6 7
	root := common.NewTreeNode(1)
	root.Left = common.NewTreeNode(2)
	root.Right = common.NewTreeNode(3)
	root.Left.Left = common.NewTreeNode(4)
	root.Left.Right = common.NewTreeNode(5)
	root.Right.Left = common.NewTreeNode(6)
	root.Right.Right = common.NewTreeNode(7)

	fmt.Println("BFS")
	b := bfs.BFS(root)
	for _, element := range b {
		fmt.Print(element.Val)
	}
	fmt.Println("RecursiveDFS")
	d := dfs.RecursiveDFS(root)
	for _, element := range d {
		fmt.Print(element.Val)
	}
	fmt.Println("IterativeDfs")
	dI := dfs.IterativeDfs(root)
	for _, element := range dI {
		fmt.Print(element.Val)
	}

	fmt.Println("Order Level Traversal")
	oL := bfs.OrderLevelTraversal(root)
	fmt.Println(oL)
}
