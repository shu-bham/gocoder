package leetcode

import (
	"fmt"
	"math"
)

func CoinChangeMinNumCoins(coins []int, amount int) int {
	dp := make([]int, amount+1)
	// Initialize dp array with a large number to represent unreachable states
	for i := range dp {
		dp[i] = math.MaxInt
	}
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		for _, coin := range coins {
			if coin <= i && dp[i-coin] != math.MaxInt {
				dp[i] = min(dp[i], 1+dp[i-coin])
			}
		}

		fmt.Println(dp[:])
	}

	// If dp[amount] is still math.MaxInt, it means the amount is unreachable
	if dp[amount] == math.MaxInt {
		return -1
	}

	return dp[amount]
}
