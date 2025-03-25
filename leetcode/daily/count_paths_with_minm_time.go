package leetcode

import (
	"container/heap"
	"fmt"
	"github.com/emirpasic/gods/v2/sets/hashset"
	"math"
)

// Medium: https://leetcode.com/problems/number-of-ways-to-arrive-at-destination

// DFS approach - Inefficient - TLE
func CountPathsWithMinimumTime(n int, roads [][]int) int {
	visited := hashset.New[int]()
	adjList := createAdjListWeighted(roads)
	path := []int{}
	path = append(path, 0)
	weight := 0
	freqMap := make(map[int]int)
	dfsHelper(0, n-1, adjList, visited, path, &weight, freqMap)
	count := 0
	smallest := math.MaxInt
	for k, v := range freqMap {
		if k < smallest {
			smallest = k
			count = v
		}
	}
	return count
}

func dfsHelper(src int, dest int, adjList map[int][][]int, visited *hashset.Set[int], path []int, pathWeight *int, freqMap map[int]int) {
	if src == dest {
		fmt.Println(path, *pathWeight)
		freqMap[*pathWeight]++
		return
	}
	visited.Add(src)
	for _, nbrs := range adjList[src] {
		if !visited.Contains(nbrs[0]) {
			path = append(path, nbrs[0])
			*pathWeight += nbrs[1]
			dfsHelper(nbrs[0], dest, adjList, visited, path, pathWeight, freqMap)
			path = path[:len(path)-1]
			*pathWeight -= nbrs[1]
		}
	}
	visited.Remove(src)
}

func createAdjListWeighted(edges [][]int) map[int][][]int {
	adjList := make(map[int][][]int)
	for _, edge := range edges {
		adjList[edge[0]] = append(adjList[edge[0]], []int{edge[1], edge[2]})
		adjList[edge[1]] = append(adjList[edge[1]], []int{edge[0], edge[2]})
	}
	return adjList
}

func CountPathsWithMinimumTimeV2(n int, roads [][]int) int {
	MOD := 1_000_000_007
	distance := make([]int, n)
	for i := 0; i < n; i++ {
		distance[i] = math.MaxInt
	}

	countOfShortestPath := make([]int, n)

	distance[0] = 0
	countOfShortestPath[0] = 1

	pq := &priorityQueue{}
	heap.Init(pq)
	heap.Push(pq, pqPair{
		distance: 0,
		vertex:   0,
	})

	adjList := createAdjListV2(roads)

	for pq.Len() > 0 {
		pop := heap.Pop(pq).(pqPair)
		u := pop.vertex

		if pop.distance > distance[u] {
			continue
		}

		for _, nbr := range adjList[u] {
			v, wt := nbr.vertex, nbr.weight

			if distance[v] > distance[u]+wt {
				distance[v] = distance[u] + wt
				heap.Push(pq, pqPair{
					distance: distance[v],
					vertex:   v,
				})
				countOfShortestPath[v] = countOfShortestPath[u]
			} else if distance[v] == distance[u]+wt {
				countOfShortestPath[v] = (countOfShortestPath[v] + countOfShortestPath[u]) % MOD
			}
		}

	}

	return countOfShortestPath[n-1]

}

func createAdjListV2(roads [][]int) map[int][]adjPair {
	m := map[int][]adjPair{}

	for _, road := range roads {
		u, v, wt := road[0], road[1], road[2]
		m[u] = append(m[u], adjPair{
			vertex: v,
			weight: wt,
		})
		m[v] = append(m[v], adjPair{
			vertex: u,
			weight: wt,
		})
	}
	return m
}

type pqPair struct {
	distance int
	vertex   int
}

type adjPair struct {
	vertex int
	weight int
}

type priorityQueue []pqPair

func (p priorityQueue) Len() int {
	return len(p)
}

func (p priorityQueue) Less(i, j int) bool {
	return p[i].distance < p[j].distance
}

func (p priorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *priorityQueue) Push(x any) {
	*p = append(*p, x.(pqPair))
}

func (p *priorityQueue) Pop() any {
	old := *p
	n := len(old)
	item := old[n-1]
	*p = old[0 : n-1]
	return item
}
