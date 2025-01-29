package leetcode

import "fmt"

func MaxProfitInfTransactionWithTransactionFee(prices []int, fee int) int {
	n := len(prices)

	if n < 2 {
		return 0
	}

	oldBuyStateProfit, oldSellStateProfit := -prices[0], 0
	fmt.Println(prices, fee)
	for i := 1; i < n; i++ {
		fmt.Println(prices[i], ":", oldBuyStateProfit, oldSellStateProfit)
		newBsp := max(oldBuyStateProfit, oldSellStateProfit-prices[i])
		newSsp := max(oldSellStateProfit, oldBuyStateProfit+prices[i]-fee)
		oldBuyStateProfit = newBsp
		oldSellStateProfit = newSsp
	}

	return oldSellStateProfit
}

func MaxProfitInfTransactionWithCoolDown(prices []int) int {
	n := len(prices)

	if n < 2 {
		return 0
	}

	oldBuyStateProfit, oldSellStateProfit, oldCoolDownStateProfit := -prices[0], 0, 0
	fmt.Println(prices)
	for i := 1; i < n; i++ {
		fmt.Println(prices[i], ":", oldBuyStateProfit, oldSellStateProfit, oldCoolDownStateProfit)
		newBsp := max(oldBuyStateProfit, oldCoolDownStateProfit-prices[i])
		newSsp := max(oldSellStateProfit, oldBuyStateProfit+prices[i])
		newCsp := max(oldCoolDownStateProfit, oldSellStateProfit)
		oldBuyStateProfit = newBsp
		oldSellStateProfit = newSsp
		oldCoolDownStateProfit = newCsp
	}

	return oldSellStateProfit
}
