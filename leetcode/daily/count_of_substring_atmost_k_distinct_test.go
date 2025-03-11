package leetcode_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	leetcode "gocoder/leetcode/daily"
)

func TestCountOfSubstringsAtmostK(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		k        int
		expected int
	}{
		{"t1", "aabacbebebe", 3, 46}, // General case
		{"t2", "abc", 2, 5},          // Small input
		{"t3", "aaaa", 1, 10},        // All same characters
		{"t4", "abaccc", 2, 15},      // Mixed characters
		{"t5", "abcdef", 3, 15},      // Unique characters
		{"t6", "abababab", 2, 36},    // Repeating pattern
		{"t7", "", 3, 0},             // Empty string
		{"t8", "abc", 0, 0},          // k = 0, no substrings possible
		{"t9", "abcde", 10, 15},      // k > distinct chars, all substrings valid
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := leetcode.CountOfSubstringsAtmostK(tt.input, tt.k)
			assert.Equal(t, tt.expected, result, "Test case failed for input: %s, k: %d", tt.input, tt.k)
		})
	}
}
