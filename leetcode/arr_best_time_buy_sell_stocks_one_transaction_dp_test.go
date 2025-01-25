package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gocoder/leetcode"
)

// Test cases for MaxProfit function
var maxProfitTestCases = []struct {
	name     string
	input    []int
	expected int
}{
	{
		name:     "Normal case with profit",
		input:    []int{7, 1, 5, 3, 6, 4},
		expected: 5,
	},
	{
		name:     "Descending prices",
		input:    []int{7, 6, 4, 3, 1},
		expected: 0,
	},
	{
		name:     "Single element array",
		input:    []int{7},
		expected: 0,
	},
	{
		name:     "Two element array with no profit",
		input:    []int{7, 6},
		expected: 0,
	},
	{
		name:     "Mixed prices with moderate profit",
		input:    []int{3, 2, 6, 5, 0, 3},
		expected: 3,
	},
}

// TestMaxProfit runs all test cases for MaxProfit function
func TestMaxProfit(t *testing.T) {
	for _, tc := range maxProfitTestCases {
		t.Run(tc.name, func(t *testing.T) {
			profit := leetcode.MaxProfit(tc.input)
			assert.Equal(t, tc.expected, profit, "Test case: %s", tc.name)
		})
	}
}
