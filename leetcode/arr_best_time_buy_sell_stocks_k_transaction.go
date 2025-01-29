package leetcode

import "fmt"

func MaxProfitKTransactions(k int, prices []int) int {
	if k < 0 {
		return 0
	}

	n := len(prices)
	// atmost k transactions
	//if n < 2*k {
	//	return 0
	//}

	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i <= k; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
				continue
			}
			temp := 0
			for p := 0; p < j; p++ {
				if prices[j] > prices[p] {
					t := dp[i-1][p] + (prices[j] - prices[p])
					temp = max(temp, t)
				}
			}
			temp = max(temp, dp[i][j-1])
			dp[i][j] = temp
			printDp(dp, k, prices)
		}
	}

	return dp[k][n-1]
}

func printDp(a [][]int, k int, p []int) {
	fmt.Println(k, p)
	fmt.Println("DP:")
	for i := range a {
		fmt.Println(a[i])
	}
	fmt.Println()
}
