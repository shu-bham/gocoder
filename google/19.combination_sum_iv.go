package google

func CombinationSum4(coins []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for amount := 1; amount <= target; amount++ {
		for _, coin := range coins {
			if coin <= amount {
				remainingAmt := amount - coin
				dp[amount] += dp[remainingAmt]
			}
		}
	}

	return dp[target]
}
