package misc

import (
	"container/heap"
	"fmt"
	"math"
)

type AdjPair struct {
	weight int
	vertex int
}

type PQPair struct {
	distance int
	vertex   int
}

type PriorityQueue []PQPair

func (p PriorityQueue) Len() int {
	return len(p)
}

func (p PriorityQueue) Less(i, j int) bool {
	return p[i].distance < p[j].distance
}

func (p PriorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PriorityQueue) Push(x any) {
	*p = append(*p, x.(PQPair))
}

func (p *PriorityQueue) Pop() any {
	old := *p
	n := len(old)
	item := old[n-1]
	*p = old[0 : n-1]
	return item
}

func addEdge(adj map[int][]AdjPair, u, v, weight int) {
	adj[u] = append(adj[u], AdjPair{weight, v})
	adj[v] = append(adj[v], AdjPair{weight, u})
}

func ShortestPathFromSrc(n int, src int, edges [][]int) []int {
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt
	}
	dist[src] = 0

	adjList := createAdjListWeighted(edges)
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, PQPair{0, src})

	for pq.Len() > 0 {
		pop := pq.Pop().(PQPair)
		u := pop.vertex

		// skip stale entries
		if pop.distance > dist[u] {
			continue
		}

		for _, nbr := range adjList[u] {
			v, weight := nbr.vertex, nbr.weight
			if dist[v] > dist[u]+weight {
				dist[v] = dist[u] + weight
				heap.Push(pq, PQPair{dist[v], v})
			}
		}
	}

	fmt.Println("Vertex \tDistance from Source")
	for i, d := range dist {
		fmt.Printf("%d \t\t %d\n", i, d)
	}

	return dist
}

func createAdjListWeighted(edges [][]int) map[int][]AdjPair {
	adj := make(map[int][]AdjPair)
	for _, edge := range edges {
		addEdge(adj, edge[0], edge[1], edge[2])
	}
	return adj
}
