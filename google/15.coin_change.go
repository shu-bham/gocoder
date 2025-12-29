package google

import "math"

// recursive
func CoinChange(coins []int, amount int) int {
	ans := minCoinRecur(0, amount, coins)
	if ans == math.MaxInt {
		return -1
	}
	return ans
}

func minCoinRecur(i int, sum int, coins []int) int {
	if sum == 0 {
		return 0
	}

	if sum < 0 || i == len(coins) {
		return math.MaxInt
	}

	// take the ith coin
	take := minCoinRecur(i, sum-coins[i], coins)
	if take != math.MaxInt {
		take++
	}

	// don't take the ith coin
	noTake := minCoinRecur(i+1, sum, coins)
	return min(take, noTake)
}
