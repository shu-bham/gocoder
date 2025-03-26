package leetcode

import (
	"math"
	"sort"
)

// Medium: https://leetcode.com/problems/minimum-operations-to-make-a-uni-value-grid
func MinOperationsToMakeUniVal(grid [][]int, x int) int {
	n, m := len(grid), len(grid[0])
	rem := grid[0][0] % x

	temp := make([]int, 0)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j]%x != rem {
				return -1
			}
			temp = append(temp, grid[i][j])
		}
	}

	sort.Ints(temp)

	// median is the optimal
	target := temp[len(temp)/2]

	targets := []int{target}

	minAns := math.MaxInt

	for _, t := range targets {
		ans := 0
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				f := float64(t - grid[i][j])
				ans += int(math.Abs(f)) / x
			}
		}
		minAns = min(minAns, ans)
	}

	return minAns

}
