package google

import (
	"testing"
)

func TestCoinChange(t *testing.T) {
	tests := []struct {
		name     string
		coins    []int
		amount   int
		expected int
	}{
		{
			name:     "standard case",
			coins:    []int{1, 2, 5},
			amount:   11,
			expected: 3, // 5 + 5 + 1
		},
		{
			name:     "no solution",
			coins:    []int{2},
			amount:   3,
			expected: -1,
		},
		{
			name:     "amount is zero",
			coins:    []int{1},
			amount:   0,
			expected: 0,
		},
		{
			name:     "single coin, exact match",
			coins:    []int{1},
			amount:   1,
			expected: 1,
		},
		{
			name:     "single coin, multiple matches",
			coins:    []int{1},
			amount:   5,
			expected: 5,
		},
		{
			name:     "larger amount, mixed coins",
			coins:    []int{1, 5, 10},
			amount:   18,
			expected: 5, // 10 + 5 + 1 + 1 + 1
		},
		{
			name:     "no coins",
			coins:    []int{},
			amount:   1,
			expected: -1,
		},
		{
			name:     "no coins, amount zero",
			coins:    []int{},
			amount:   0,
			expected: 0,
		},
		{
			name:     "coins [3,5], amount 7, no solution",
			coins:    []int{3, 5},
			amount:   7,
			expected: -1,
		},
		{
			name:     "coins [2], amount 1, no solution",
			coins:    []int{2},
			amount:   1,
			expected: -1,
		},
		{
			name:     "coins [1], amount 1, solution",
			coins:    []int{1},
			amount:   1,
			expected: 1,
		},
		{
			name:     "tle",
			coins:    []int{1, 2, 5},
			amount:   100,
			expected: 20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CoinChange(tt.coins, tt.amount)
			if got != tt.expected {
				t.Errorf("CoinChange(%v, %d) got = %d, want %d", tt.coins, tt.amount, got, tt.expected)
			}
		})
	}
}
