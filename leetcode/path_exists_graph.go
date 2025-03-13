package leetcode

type graph struct {
	adjList map[int][]int
}

func newGraph() *graph {
	return &graph{adjList: make(map[int][]int)}
}

func (g *graph) addEdge(u, v int) {
	g.adjList[u] = append(g.adjList[u], v)
	g.adjList[v] = append(g.adjList[v], u)
}

func (g *graph) getAdjList() map[int][]int {
	return g.adjList
}

func ValidPath(n int, edges [][]int, source int, destination int) bool {
	g := createGraph(edges)
	visited := make(map[int]bool)
	queue := []int{source}
	visited[source] = true

	for len(queue) > 0 {
		curr := queue[0]
		if curr == destination {
			return true
		}
		queue = queue[1:]
		for _, nbr := range g.getAdjList()[curr] {
			if !visited[nbr] {
				visited[nbr] = true
				queue = append(queue, nbr)
			}
		}

	}

	return false
}

func createGraph(edges [][]int) *graph {
	g := newGraph()
	for _, edge := range edges {
		g.addEdge(edge[0], edge[1])
	}
	return g
}
