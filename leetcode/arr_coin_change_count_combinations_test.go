package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/leetcode"
	"testing"
)

func TestCoinChangeCountCombinations(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		coins := []int{1, 2, 5}
		amount := 5
		combinations := leetcode.CoinChangeCountCombinations(coins, amount)
		assert.Equal(t, 4, combinations)

	})

	t.Run("t2", func(t *testing.T) {
		coins := []int{1}
		amount := 5
		combinations := leetcode.CoinChangeCountCombinations(coins, amount)
		assert.Equal(t, 1, combinations)

	})

	t.Run("t3", func(t *testing.T) {
		coins := []int{1, 3}
		amount := 4
		combinations := leetcode.CoinChangeCountCombinations(coins, amount)
		assert.Equal(t, 2, combinations)

	})

	t.Run("t4", func(t *testing.T) {
		coins := []int{2, 5}
		amount := 10
		combinations := leetcode.CoinChangeCountCombinations(coins, amount)
		assert.Equal(t, 2, combinations)

	})
}

func TestCoinChangeCountCombinationsV2(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		coins := []int{1, 2, 5}
		amount := 5
		combinations := leetcode.CoinChangeCountCombinationsV2(coins, amount)
		assert.Equal(t, 4, combinations)

	})

	t.Run("t2", func(t *testing.T) {
		coins := []int{1}
		amount := 5
		combinations := leetcode.CoinChangeCountCombinationsV2(coins, amount)
		assert.Equal(t, 1, combinations)

	})

	t.Run("t3", func(t *testing.T) {
		coins := []int{1, 3}
		amount := 4
		combinations := leetcode.CoinChangeCountCombinationsV2(coins, amount)
		assert.Equal(t, 2, combinations)

	})

	t.Run("t4", func(t *testing.T) {
		coins := []int{2, 5}
		amount := 10
		combinations := leetcode.CoinChangeCountCombinationsV2(coins, amount)
		assert.Equal(t, 2, combinations)

	})
}
