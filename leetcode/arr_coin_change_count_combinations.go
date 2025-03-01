package leetcode

func CoinChangeCountCombinations(coins []int, amount int) int {
	row, col := len(coins)+1, amount+1
	dp := make([][]int, row)
	for i := range dp {
		dp[i] = make([]int, col)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = 1
			} else if i == 0 {
				dp[i][j] = 0
			} else if j == 0 {
				dp[i][j] = 1
			} else {
				coin := coins[i-1]
				dp[i][j] = dp[i-1][j]
				if j >= coin {
					dp[i][j] += dp[i][j-coin]
				}
			}
		}
	}

	return dp[row-1][amount]
}

// intuition : since in prev approach we are effectively copying
// prev row to curr row and then doing processing
// we can achieve this using a 1D array and switching the inner and outer iterations
func CoinChangeCountCombinationsV2(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for i := 1; i <= amount; i++ {
			if i >= coin {
				dp[i] += dp[i-coin]
			}
		}
	}

	return dp[amount]
}
