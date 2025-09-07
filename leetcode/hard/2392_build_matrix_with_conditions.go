package hard

import "errors"

/**
https://leetcode.com/problems/build-a-matrix-with-conditions

You are given a positive integer k. You are also given:

a 2D integer array rowConditions of size n where rowConditions[i] = [abovei, belowi], and
a 2D integer array colConditions of size m where colConditions[i] = [lefti, righti].
The two arrays contain integers from 1 to k.

You have to build a k x k matrix that contains each of the numbers from 1 to k exactly once. The remaining cells should have the value 0.

The matrix should also satisfy the following conditions:

The number abovei should appear in a row that is strictly above the row at which the number belowi appears for all i from 0 to n - 1.
The number lefti should appear in a column that is strictly left of the column at which the number righti appears for all i from 0 to m - 1.
Return any matrix that satisfies the conditions. If no answer exists, return an empty matrix.
*/

func BuildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	emptyRes := [][]int{}
	// get topological sort of rowCondition to fetch the desired order
	rowOrder, err := getTopologicalOrder(rowConditions, k)
	if err != nil {
		// when cycle is present
		return emptyRes
	}
	// get topological sort of colConditions to fetch the desired order
	colOrder, err := getTopologicalOrder(colConditions, k)
	if err != nil {
		return emptyRes
	}

	res := make([][]int, k)

	// rowOrder and colOrder will have k elements
	// generate the matrix from rowOrder and colOrder
	for i := 0; i < k; i++ {
		res[i] = make([]int, k)
		for j := 0; j < k; j++ {
			if rowOrder[i] == colOrder[j] {
				res[i][j] = rowOrder[i]
				break
			}
		}
	}

	return res
}

// use Kahn's algo to find topological sort as it can also help in cycle detection
func getTopologicalOrder(edges [][]int, k int) ([]int, error) {
	// create adjList
	adjList := make(map[int][]int)
	for _, edge := range edges {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
	}

	indegree := make([]int, k+1)
	for _, edge := range edges {
		indegree[edge[1]]++
	}

	// enqueue all nodes with indegree of 0
	var queue []int
	for i := 1; i <= k; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	var res []int

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		res = append(res, front)

		// decrease indegree of all nbr's and enqueue one with indegree=0
		for _, nbr := range adjList[front] {
			indegree[nbr]--
			if indegree[nbr] == 0 {
				queue = append(queue, nbr)
			}
		}

	}

	// if all nodes were not reached means cycle is present
	if len(res) != k {
		return nil, errors.New("cycle")
	}

	return res, nil
}
