package main

import (
	"fmt"
	"github.com/aaronangxz/TrainingGround/common"
	"github.com/aaronangxz/TrainingGround/graph/bfs"
)

func main() {
	g := common.NewGraph(5)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 3)

	fmt.Println(bfs.BFS(g, 2))

	fmt.Println(bfs.BFSAllVertices(g))
}
