package common

type Graph struct {
	Vertices      int
	AdjacencyList [][]int
}

func NewGraph(vertices int) *Graph {
	return &Graph{
		Vertices:      vertices,
		AdjacencyList: make([][]int, vertices),
	}
}

func (g *Graph) AddEdge(v, w int) {
	g.AdjacencyList[v] = append(g.AdjacencyList[v], w)
}
