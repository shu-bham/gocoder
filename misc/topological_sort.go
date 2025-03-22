package misc

import (
	"container/list"
	"github.com/emirpasic/gods/v2/sets/hashset"
)

func TopologicalSortInDegree(n int, edges [][]int) ([]int, bool) {
	indegree := make([]int, n)
	for _, edge := range edges {
		indegree[edge[1]]++
	}

	queue := list.New()
	for i, deg := range indegree {
		if deg == 0 {
			queue.PushBack(i)
		}
	}

	adjList := createAdjList(edges)

	var ans []int
	for queue.Len() > 0 {
		curr := queue.Remove(queue.Front()).(int) // Dequeue
		ans = append(ans, curr)
		for _, nbr := range adjList[curr] {
			indegree[nbr]--
			if indegree[nbr] == 0 {
				queue.PushBack(nbr) // Enqueue
			}
		}
	}

	if len(ans) != n {
		return nil, false
	}

	return ans, true
}

// only for DAG, can't detect cycle
func TopologicalSortDFS(n int, edges [][]int) ([]int, bool) {
	visited := hashset.New[int]()
	stack := make([]int, 0)

	adjList := createAdjList(edges)
	for i := 0; i < n; i++ {
		if !visited.Contains(i) {
			dfsUtil(i, adjList, visited, &stack)
		}
	}

	var ans []int
	for i := len(stack) - 1; i >= 0; i-- {
		ans = append(ans, stack[i])
	}
	return ans, true

}

func dfsUtil(node int, adjList map[int][]int, visited *hashset.Set[int], stack *[]int) {
	visited.Add(node)
	for _, nbr := range adjList[node] {
		if !visited.Contains(nbr) {
			dfsUtil(nbr, adjList, visited, stack)
		}
	}
	*stack = append(*stack, node)
}

func createAdjList(edges [][]int) map[int][]int {
	adjListMap := make(map[int][]int)
	for _, edge := range edges {
		adjListMap[edge[0]] = append(adjListMap[edge[0]], edge[1])
	}

	return adjListMap
}
