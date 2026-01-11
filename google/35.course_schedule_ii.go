package google

func FindOrderDFS(numCourses int, prereq [][]int) []int {
	adjList := make(map[int][]int)
	visited := make([]int, numCourses)
	var stack []int

	for _, edge := range prereq {
		list, ok := adjList[edge[0]]
		if ok {
			list = append(list, edge[1])
			adjList[edge[0]] = list
		} else {
			l := []int{edge[1]}
			adjList[edge[0]] = l
		}
	}

	// fmt.Println(adjList)

	var dfsUtil func(int) bool
	dfsUtil = func(node int) bool {
		if visited[node] == 1 { // Cycle detected
			return false
		}
		if visited[node] == 2 { // Already processed
			return true
		}

		visited[node] = 1 // Mark as visiting

		for _, nbr := range adjList[node] {
			if !dfsUtil(nbr) {
				return false
			}
		}

		visited[node] = 2 // Mark as visited
		stack = append(stack, node)
		return true
	}

	for i := 0; i < numCourses; i++ {
		if visited[i] == 0 {
			if !dfsUtil(i) {
				return []int{} // Cycle detected
			}
		}
	}

	return stack
}

func FindOrderIndegreeBFS(numCourses int, prereq [][]int) []int {
	adjList := make(map[int][]int)
	indegree := make([]int, numCourses)

	for _, edge := range prereq {
		indegree[edge[1]]++
		list, ok := adjList[edge[0]]
		if ok {
			list = append(list, edge[1])
			adjList[edge[0]] = list
		} else {
			l := []int{edge[1]}
			adjList[edge[0]] = l
		}
	}

	que := make([]int, 0)
	for i, val := range indegree {
		if val == 0 {
			que = append(que, i)
		}
	}

	var ans []int
	for len(que) > 0 {
		front := que[0]
		que = que[1:]
		ans = append(ans, front)

		for _, nbr := range adjList[front] {
			indegree[nbr]--
			if indegree[nbr] == 0 {
				que = append(que, nbr)
			}
		}

	}

	if len(ans) != numCourses {
		return []int{}
	}

	reverse(ans)
	return ans
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
