package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/leetcode"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	profit := leetcode.MaxProfit([]int{7, 1, 5, 3, 6, 4})
	assert.Equal(t, 5, profit)
}

func TestMaxProfit2(t *testing.T) {
	profit := leetcode.MaxProfit([]int{7, 6, 4, 3, 1})
	assert.Equal(t, 0, profit)
}

func TestMaxProfit3(t *testing.T) {
	profit := leetcode.MaxProfit([]int{7})
	assert.Equal(t, 0, profit)
}

func TestMaxProfit4(t *testing.T) {
	profit := leetcode.MaxProfit([]int{7, 6})
	assert.Equal(t, 0, profit)
}

func TestMaxProfit5(t *testing.T) {
	profit := leetcode.MaxProfit([]int{3, 2, 6, 5, 0, 3})
	assert.Equal(t, 4, profit)
}
