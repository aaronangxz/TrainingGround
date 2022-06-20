package bfs

import "github.com/aaronangxz/TrainingGround/common"

//BFS Traverse reachable vertices
func BFS(g *common.Graph, s int) []int {
	var (
		//keep track of visited vertices, all false
		visited = make([]bool, g.Vertices)
		//queue to keep track of the next vertex to visit
		q   common.Queue
		out []int
	)

	//visit the first vertex, put in queue
	visited[s] = true
	q.Enqueue(s)

	//Continue as long as queue is not empty
	for !q.IsEmpty() {
		//visit the first vertex in queue
		s = q.Front().(int)
		out = append(out, s)
		//pop the first vertex (visited)
		q.Dequeue()

		//find the adjacent vertices of this vertex
		for _, a := range g.AdjacencyList[s] {
			//If they are not visited, mark as visited and put into queue
			if !visited[a] {
				visited[a] = true
				q.Enqueue(a)
			}
		}
	}
	return out
}

func BFSAllVertices(g *common.Graph) []int {
	var (
		//keep track of visited vertices, all false
		visited = make([]bool, g.Vertices)
		//queue to keep track of the next vertex to visit
		out []int
	)
	for i := 0; i < g.Vertices; i++ {
		if !visited[i] {
			var q common.Queue
			visited[i] = true
			q.Enqueue(i)

			for !q.IsEmpty() {
				v := q.Front().(int)
				q.Dequeue()
				out = append(out, v)

				for _, a := range g.AdjacencyList[v] {
					if !visited[a] {
						visited[a] = true
						q.Enqueue(a)
					}
				}
			}
		}
	}
	return out
}
