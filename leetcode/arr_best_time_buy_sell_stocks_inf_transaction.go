package leetcode

func MaxProfitInfTransaction(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	totalProfit := 0
	buyPos, sellPos := 0, 0
	for i := 1; i < len(prices); i++ {
		if prices[i] >= prices[i-1] {
			sellPos++
		} else {
			totalProfit += prices[sellPos] - prices[buyPos]
			buyPos, sellPos = i, i
		}
	}
	totalProfit += prices[sellPos] - prices[buyPos]

	return totalProfit
}
