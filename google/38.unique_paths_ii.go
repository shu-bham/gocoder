package google

func UniquePathsWithObstacles(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	if grid[m-1][n-1] == 1 {
		return 0
	}

	const MOD = 2_000_000_000

	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	ans[m-1][n-1] = 1

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] == 1 {
				continue
			}
			if i+1 < m {
				ans[i][j] += (ans[i+1][j] % MOD)
			}
			if j+1 < n {
				ans[i][j] += (ans[i][j+1] % MOD)
			}
		}
	}

	return ans[0][0]

}
