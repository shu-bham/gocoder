package misc

import "fmt"

type Graph struct {
	adjList map[int][]int
}

func NewGraph() *Graph {
	return &Graph{adjList: make(map[int][]int)}
}

func (g *Graph) AddEdge(u, v int) {
	g.adjList[u] = append(g.adjList[u], v)
	g.adjList[v] = append(g.adjList[v], u)
}

func (g *Graph) GetAdjList() map[int][]int {
	return g.adjList
}

func BFS(edges [][]int, src int) []int {
	g := makeGraph(edges)
	visited := make(map[int]bool)
	queue := []int{src}
	visited[src] = true
	var res []int
	for len(queue) > 0 {
		curr := queue[0]
		fmt.Print(curr, " ")
		res = append(res, curr)
		queue = queue[1:]

		for _, nbr := range g.GetAdjList()[curr] {
			if !visited[nbr] {
				visited[nbr] = true
				queue = append(queue, nbr)
			}
		}
	}
	return res
}

func makeGraph(edges [][]int) *Graph {
	graph := NewGraph()
	for _, edge := range edges {
		graph.AddEdge(edge[0], edge[1])
	}
	return graph
}
