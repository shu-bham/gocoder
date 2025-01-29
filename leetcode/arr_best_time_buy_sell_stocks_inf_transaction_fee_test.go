package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gocoder/leetcode"
)

// Test cases for MaxProfitInfTransaction function
var MaxProfitInfTransactionWithTransactionFee = []struct {
	name     string
	input    []int
	fee      int
	expected int
}{
	{
		name:     "Normal case with profit",
		input:    []int{7, 1, 5, 3, 6, 4},
		fee:      0,
		expected: 7,
	},
	{
		name:     "Descending prices",
		input:    []int{7, 6, 4, 3, 1},
		fee:      0,
		expected: 0,
	},
	{
		name:     "Single element array",
		input:    []int{7},
		fee:      0,
		expected: 0,
	},
	{
		name:     "Two element array with no profit",
		input:    []int{7, 6},
		fee:      0,
		expected: 0,
	},
	{
		name:     "Mixed prices with moderate profit",
		input:    []int{3, 2, 6, 5, 0, 3},
		fee:      0,
		expected: 7,
	},
	{
		name:     "Mixed prices with moderate profit v2",
		input:    []int{10, 15, 17, 20, 16, 18, 22, 20, 22, 20, 23, 25},
		fee:      3,
		expected: 13,
	},
}

// MaxProfitInfTransactionWithTransactionFee runs all test cases for MaxProfitInfTransaction function
func TestMaxProfitInfTransactionWithTransactionFee(t *testing.T) {
	for _, tc := range MaxProfitInfTransactionWithTransactionFee {
		t.Run(tc.name, func(t *testing.T) {
			profit := leetcode.MaxProfitInfTransactionWithTransactionFee(tc.input, tc.fee)
			assert.Equal(t, tc.expected, profit, "Test case: %s", tc.name)
		})
	}
}
