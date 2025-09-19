package medium

import (
	"fmt"
)

/**
62. Unique Paths
https://leetcode.com/problems/unique-paths/
There is a robot on an m x n grid. The robot is initially located at the top-left corner (i.e., grid[0][0]). The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]). The robot can only move either down or right at any point in time.

Given the two integers m and n, return the number of possible unique paths that the robot can take to reach the bottom-right corner.

The test cases are generated so that the answer will be less than or equal to 2 * 109.
*/

const MOD = 2_000_000_000

type pair struct {
	i, j int
}

// DP - Tabulation
func UniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// dp[i][j] will denote the number of unique paths from cell (0, 0) the cell (i, j)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = ((dp[i-1][j] % MOD) + (dp[i][j-1] % MOD)) % MOD
			}
		}
	}

	return dp[m-1][n-1]
}

// DP - Memoisation
func UniquePathsV2(m int, n int) int {
	dp := make(map[pair]int)
	// dp[i][j] denotes number of unique paths from cell (i, j) to the cell (m-1,n-1)
	dp[pair{m - 1, n - 1}] = 1
	return calcWays(0, 0, m, n, dp)
}

func calcWays(i, j, m, n int, dp map[pair]int) int {
	if i == m || j == n {
		return 0
	}

	res, ok := dp[pair{i, j}]
	if ok {
		return res
	}

	dp[pair{i + 1, j}] = calcWays(i+1, j, m, n, dp) % MOD
	dp[pair{i, j + 1}] = calcWays(i, j+1, m, n, dp) % MOD
	dp[pair{i, j}] = (dp[pair{i + 1, j}] + dp[pair{i, j + 1}]) % MOD
	return dp[pair{i, j}]
}

// TLE, 31/63 TCs, Use of Combinatorics
func UniquePathsSlow(m int, n int) int {
	dp := make(map[int]int)
	dp[0] = 1
	dp[1] = 1

	f1 := fact(m-1+n-1, dp)
	f2 := fact(m-1, dp)
	f3 := fact(n-1, dp)

	fmt.Println(f1, f2, f3)

	return f1 % MOD / (f2 * f3)

}

func fact(n int, dp map[int]int) int {
	res, ok := dp[n]
	if ok {
		return res
	}

	dp[n] = n * fact(n-1, dp)
	return dp[n]
}
