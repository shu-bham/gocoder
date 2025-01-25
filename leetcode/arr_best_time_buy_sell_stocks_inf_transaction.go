package leetcode

func MaxProfitInfTransaction(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	totalProfit := 0
	buy_pos, sell_pos := 0, 0
	for i := 1; i < len(prices); i++ {
		if prices[i] >= prices[i-1] {
			sell_pos++
		} else {
			totalProfit += prices[sell_pos] - prices[buy_pos]
			buy_pos, sell_pos = i, i
		}
	}
	totalProfit += prices[sell_pos] - prices[buy_pos]

	return totalProfit
}
