package google_test

import (
	"gocoder/google"
	"testing"
)

func TestCoinChangeII(t *testing.T) {
	tests := []struct {
		name     string
		amount   int
		coins    []int
		expected int
	}{
		{
			name:     "standard case - amount 5, coins [1,2,5]",
			amount:   5,
			coins:    []int{1, 2, 5},
			expected: 4, // 5, 2+2+1, 2+1+1+1, 1+1+1+1+1
		},
		{
			name:     "amount 3, coins [1,2]",
			amount:   3,
			coins:    []int{1, 2},
			expected: 2, // 1+1+1, 1+2
		},
		{
			name:     "no coins",
			amount:   3,
			coins:    []int{},
			expected: 0,
		},
		{
			name:     "amount is zero",
			amount:   0,
			coins:    []int{1, 2, 5},
			expected: 1, // one way to make 0: use no coins
		},
		{
			name:     "no solution",
			amount:   3,
			coins:    []int{2},
			expected: 0,
		},
		{
			name:     "only one way",
			amount:   1,
			coins:    []int{1},
			expected: 1,
		},
		{
			name:     "different coins, exact match",
			amount:   10,
			coins:    []int{10},
			expected: 1,
		},
		{
			name:     "amount 5, coins [3]",
			amount:   5,
			coins:    []int{3},
			expected: 0,
		},
		{
			name:     "amount 10, coins [3]",
			amount:   10,
			coins:    []int{3},
			expected: 0,
		},
		{
			name:     "amount 4, coins [1,2,3]",
			amount:   4,
			coins:    []int{1, 2, 3},
			expected: 4, // {1,1,1,1}, {1,1,2}, {1,3}, {2,2}
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := google.CoinChangeII(tt.amount, tt.coins)
			if got != tt.expected {
				t.Errorf("CoinChangeII(%d, %v) got = %d, want %d", tt.amount, tt.coins, got, tt.expected)
			}
		})
	}
}

func TestCoinChangeIIv2(t *testing.T) {
	tests := []struct {
		name     string
		amount   int
		coins    []int
		expected int
	}{
		{
			name:     "standard case - amount 5, coins [1,2,5]",
			amount:   5,
			coins:    []int{1, 2, 5},
			expected: 4, // 5, 2+2+1, 2+1+1+1, 1+1+1+1+1
		},
		{
			name:     "amount 3, coins [1,2]",
			amount:   3,
			coins:    []int{1, 2},
			expected: 2, // 1+1+1, 1+2
		},
		{
			name:     "no coins",
			amount:   3,
			coins:    []int{},
			expected: 0,
		},
		{
			name:     "amount is zero",
			amount:   0,
			coins:    []int{1, 2, 5},
			expected: 1, // one way to make 0: use no coins
		},
		{
			name:     "no solution",
			amount:   3,
			coins:    []int{2},
			expected: 0,
		},
		{
			name:     "only one way",
			amount:   1,
			coins:    []int{1},
			expected: 1,
		},
		{
			name:     "different coins, exact match",
			amount:   10,
			coins:    []int{10},
			expected: 1,
		},
		{
			name:     "amount 5, coins [3]",
			amount:   5,
			coins:    []int{3},
			expected: 0,
		},
		{
			name:     "amount 10, coins [3]",
			amount:   10,
			coins:    []int{3},
			expected: 0,
		},
		{
			name:     "amount 4, coins [1,2,3]",
			amount:   4,
			coins:    []int{1, 2, 3},
			expected: 4, // {1,1,1,1}, {1,1,2}, {1,3}, {2,2}
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := google.CoinChangeIIv2(tt.amount, tt.coins)
			if got != tt.expected {
				t.Errorf("CoinChangeII(%d, %v) got = %d, want %d", tt.amount, tt.coins, got, tt.expected)
			}
		})
	}

	google.IsPalindrome("race a car")
}
