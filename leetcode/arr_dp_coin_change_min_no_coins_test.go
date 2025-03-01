package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/leetcode"
	"testing"
)

func TestCoinChangeMinNumCoins(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		coins := []int{1, 3, 5}
		amount := 7
		numCoins := leetcode.CoinChangeMinNumCoins(coins, amount)
		assert.Equal(t, 3, numCoins)
	})

	t.Run("t2", func(t *testing.T) {
		coins := []int{5}
		amount := 3
		numCoins := leetcode.CoinChangeMinNumCoins(coins, amount)
		assert.Equal(t, -1, numCoins)
	})

	t.Run("t3", func(t *testing.T) {
		coins := []int{1, 2, 5}
		amount := 17
		numCoins := leetcode.CoinChangeMinNumCoins(coins, amount)
		assert.Equal(t, 4, numCoins)
	})

	t.Run("t4", func(t *testing.T) {
		coins := []int{1, 5, 7}
		amount := 100
		numCoins := leetcode.CoinChangeMinNumCoins(coins, amount)
		assert.Equal(t, 16, numCoins)
	})

	t.Run("t5", func(t *testing.T) {
		coins := []int{3, 10, 7}
		amount := 200
		numCoins := leetcode.CoinChangeMinNumCoins(coins, amount)
		assert.Equal(t, 20, numCoins)
	})

	t.Run("t6", func(t *testing.T) {
		coins := []int{3}
		amount := 4
		numCoins := leetcode.CoinChangeMinNumCoins(coins, amount)
		assert.Equal(t, -1, numCoins)
	})
}
