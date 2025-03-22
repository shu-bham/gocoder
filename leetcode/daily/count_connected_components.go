package leetcode

// Medium: https://leetcode.com/problems/count-the-number-of-complete-components
func CountCompleteComponents(n int, edges [][]int) int {
	visited := make(map[int]bool)
	adjList := createAdjList(edges)

	ans := 0
	for i := 0; i < n; i++ {
		edgeCount := 0
		nodeCount := 0
		if !visited[i] {
			dfs(i, visited, adjList, &edgeCount, &nodeCount)
			edgeCount = edgeCount / 2 // doubly counted
			if edgeCount == nodeCount*(nodeCount-1)/2 {
				ans++
			}
		}
	}

	return ans
}

func dfs(node int, visited map[int]bool, adjList map[int][]int, edgeCount *int, nodeCount *int) {
	visited[node] = true
	*nodeCount++
	*edgeCount = (*edgeCount) + len(adjList[node])
	for _, nbr := range adjList[node] {
		if !visited[nbr] {
			dfs(nbr, visited, adjList, edgeCount, nodeCount)
		}
	}
}

func createAdjList(edges [][]int) map[int][]int {
	adjList := make(map[int][]int)
	for _, edge := range edges {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}
	return adjList
}
