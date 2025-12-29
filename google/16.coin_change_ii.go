package google

// recursive - count no. of ways to have amount as sum
func CoinChangeII(amount int, coins []int) int {
	return coinChangeIIrecur(amount, coins, len(coins))
}

func coinChangeIIrecur(sum int, coins []int, n int) int {
	if sum == 0 {
		return 1
	}

	if sum < 0 || n == 0 {
		return 0
	}
	return coinChangeIIrecur(sum, coins, n-1) + coinChangeIIrecur(sum-coins[n-1], coins, n)
}

// dp - bottom up
func CoinChangeIIv2(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			dp[j] += dp[j-coin]
		}
	}

	return dp[amount]
}
