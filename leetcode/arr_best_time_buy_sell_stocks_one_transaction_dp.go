package leetcode

// Best Time to Buy and Sell Stocks - One Transaction Allowed - Dynamic Programming

// bruteforce
func MaxProfitBF(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}

	maxPr := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if prices[j] > prices[i] {
				maxPr = max(maxPr, prices[j]-prices[i])
			}
		}
	}

	return maxPr
}

// optimised
func MaxProfit(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}

	maxPr := 0
	minSoFar := prices[0]

	for i := 1; i < n; i++ {
		if minSoFar < prices[i] {
			profitIfSoldToday := prices[i] - minSoFar
			maxPr = max(maxPr, profitIfSoldToday)
		} else {
			minSoFar = min(minSoFar, prices[i])
		}

	}

	return maxPr
}
