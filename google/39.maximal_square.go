package google

import "math"

func MaximalSquare(grid [][]byte) int {
	m, n := len(grid), len(grid[0])

	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}

	ans := 0
	dir := [][]int{{0, 1}, {1, 0}, {1, 1}}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 || j == n-1 {
				if grid[i][j] == '1' {
					ans = max(ans, 1)
					res[i][j] = 1
				}
			} else if grid[i][j] == '1' {
				minLen := math.MaxInt
				for _, point := range dir {
					x := i + point[0]
					y := j + point[1]
					minLen = min(minLen, res[x][y])
				}
				res[i][j] = 1 + minLen
				ans = max(ans, res[i][j])
			}
		}
	}
	return ans * ans
}
