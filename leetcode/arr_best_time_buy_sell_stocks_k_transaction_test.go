package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gocoder/leetcode"
)

// Test cases for maxProfitKTransactionTestCases function
var maxProfitKTransactionTestCases = []struct {
	name     string
	k        int
	input    []int
	expected int
}{
	{
		name:     "Normal case with profit",
		k:        3,
		input:    []int{7, 1, 5, 3, 6, 4, 10},
		expected: 13,
	},
	{
		name:     "Descending prices",
		k:        1,
		input:    []int{7, 6, 4, 3, 1},
		expected: 0,
	},
	{
		name:     "Single element array",
		k:        1,
		input:    []int{7},
		expected: 0,
	},
	{
		name:     "Two element array with no profit",
		k:        1,
		input:    []int{7, 6},
		expected: 0,
	},
	{
		name:     "Mixed prices with moderate profit",
		k:        3,
		input:    []int{3, 2, 6, 5, 0, 3, 8, 14},
		expected: 18,
	},
}

// TestMaxProfitInfTransaction runs all test cases for MaxProfitInfTransaction function
func TestMaxProfitKTransaction(t *testing.T) {
	for _, tc := range maxProfitKTransactionTestCases {
		t.Run(tc.name, func(t *testing.T) {
			profit := leetcode.MaxProfitKTransactions(tc.k, tc.input)
			assert.Equal(t, tc.expected, profit, "Test case: %s", tc.name)
		})
	}
}
