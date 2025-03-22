package leetcode

import "math"

// Hard: https://leetcode.com/problems/minimum-cost-walk-in-weighted-graph

func MinimumCostWalkInWeightedGraph(n int, edges [][]int, query [][]int) []int {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = -1
	}

	depth := make([]int, n)
	componentCost := make([]int, n)

	for i := range componentCost {
		componentCost[i] = math.MaxInt
	}

	for _, edge := range edges {
		union(depth, parent, edge[0], edge[1])
	}

	for _, edge := range edges {
		root := find(parent, edge[0])
		componentCost[root] = componentCost[root] & edge[2]
	}

	answer := make([]int, len(query))
	for i, q := range query {
		start := q[0]
		end := q[1]

		if find(parent, start) != find(parent, end) {
			answer[i] = -1
		} else {
			root := find(parent, start)
			answer[i] = componentCost[root]
		}
	}

	return answer
}

func union(depth []int, parent []int, i int, j int) {
	root1 := find(parent, i)
	root2 := find(parent, j)
	if root1 == root2 {
		return
	}

	if depth[root1] < depth[root2] {
		root1, root2 = root2, root1
	}

	parent[root2] = root1

	if depth[root1] == depth[root2] {
		depth[root1]++
	}

}

func find(parent []int, node int) int {
	if parent[node] == -1 {
		return node
	}

	parent[node] = find(parent, parent[node])
	return parent[node]
}
