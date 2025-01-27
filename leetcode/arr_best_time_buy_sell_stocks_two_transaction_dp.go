package leetcode

import (
	"math"
)

func MaxProfitTwoTransaction(prices []int) int {
	if len(prices) < 4 {
		return 0
	}

	//fmt.Println(prices)

	minSoFar := math.MaxInt
	profitIfSoldTodayOrBefore := make([]int, len(prices))
	for i, val := range prices {
		if i == 0 {
			minSoFar = min(minSoFar, val)
			profitIfSoldTodayOrBefore[i] = 0
			continue
		}

		profitIfSoldToday := max(0, val-minSoFar)
		profitIfSoldTodayOrBefore[i] = max(profitIfSoldTodayOrBefore[i-1], profitIfSoldToday)
		minSoFar = min(minSoFar, val)
		//fmt.Println(profitIfSoldToday, minSoFar, profitIfSoldTodayOrBefore)
	}

	maxSoFar := math.MinInt
	profitIfBuyTodayOrAfter := make([]int, len(prices))
	for i := len(prices) - 1; i >= 0; i-- {
		val := prices[i]
		if i == len(prices)-1 {
			maxSoFar = max(maxSoFar, val)
			profitIfBuyTodayOrAfter[i] = 0
			continue
		}

		profitIfBuyToday := max(0, maxSoFar-val)
		profitIfBuyTodayOrAfter[i] = max(profitIfBuyTodayOrAfter[i+1], profitIfBuyToday)
		maxSoFar = max(maxSoFar, val)
		//fmt.Println(profitIfBuyToday, maxSoFar, profitIfBuyTodayOrAfter)
	}

	//fmt.Println(profitIfSoldTodayOrBefore, profitIfBuyTodayOrAfter)
	maxProfitWithExactTwoTransactions := 0
	for i := 0; i < len(prices); i++ {
		maxProfitWithExactTwoTransactions = max(maxProfitWithExactTwoTransactions, profitIfSoldTodayOrBefore[i]+profitIfBuyTodayOrAfter[i])
	}

	return maxProfitWithExactTwoTransactions
}
