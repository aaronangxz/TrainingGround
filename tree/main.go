package main

import (
	"fmt"

	"github.com/aaronangxz/TrainingGround/common"
	"github.com/aaronangxz/TrainingGround/tree/bfs"
	"github.com/aaronangxz/TrainingGround/tree/dfs"
)

func main() {
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

	b := bfs.BFS(root)
	for _, element := range b {
		fmt.Print(element.Val)
	}
	fmt.Println()
	d := dfs.RecursiveDFS(root)
	for _, element := range d {
		fmt.Print(element.Val)
	}
	fmt.Println()
	dI := dfs.IterativeDfs(root)
	for _, element := range dI {
		fmt.Print(element.Val)
	}
}
