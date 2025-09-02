package leetcode

import (
	"fmt"
	"github.com/emirpasic/gods/v2/sets/hashset"
)

// Hard: https://leetcode.com/problems/maximum-number-of-points-from-grid-queries

// BruteForce DFS
func MaxPointsFromGridQueries(grid [][]int, queries []int) []int {

	ans := make([]int, len(queries))

	for i, query := range queries {
		visited := hashset.New[string]()
		tempAns, ok := dfsMaxPointHelper(0, 0, grid, query, visited)
		if ok {
			ans[i] = tempAns
		} else {
			ans[i] = 0
		}
	}

	return ans

}

func dfsMaxPointHelper(i int, j int, grid [][]int, query int, visited *hashset.Set[string]) (int, bool) {
	key := fmt.Sprintf("%d_%d", i, j)
	m, n := len(grid), len(grid[0])
	if i < 0 || j < 0 || i >= m || j >= n || visited.Contains(key) || grid[i][j] >= query {
		return 0, false
	}

	visited.Add(key)
	ans := 0
	all := [][]int{{i, j + 1}, {i + 1, j}}
	for _, xy := range all {
		next := fmt.Sprintf("%d_%d", xy[0], xy[1])
		if visited.Contains(next) {
			continue
		}
		intVal, ok := dfsMaxPointHelper(xy[0], xy[1], grid, query, visited)
		if ok {
			ans += intVal
		}
	}

	return ans + 1, true
}
