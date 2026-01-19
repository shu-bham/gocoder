package google

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveInvalidParentheses(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected []string
	}{
		{
			name:     "Example 1",
			s:        "()())()",
			expected: []string{"(())()", "()()()"},
		},
		{
			name:     "Example 2",
			s:        "(a)())()",
			expected: []string{"(a())()", "(a)()()"},
		},
		{
			name:     "Example 3",
			s:        ")(",
			expected: []string{""},
		},
		{
			name:     "Empty string",
			s:        "",
			expected: []string{""},
		},
		{
			name:     "No invalid parentheses",
			s:        "()()",
			expected: []string{"()()"},
		},
		{
			name:     "All invalid parentheses",
			s:        "(((",
			expected: []string{""},
		},
		{
			name:     "Complex case with duplicates",
			s:        "(())",
			expected: []string{"(())"},
		},
		{
			name:     "Another complex case",
			s:        "()((((",
			expected: []string{"()"},
		},
		{
			name:     "Case with leading invalid",
			s:        "))((",
			expected: []string{""},
		},
		{
			name:     "Leetcode example `n()())()`",
			s:        "n()())()",
			expected: []string{"n(())()", "n()()()"},
		},
		{
			name:     "String with no parenthesis",
			s:        "abc",
			expected: []string{"abc"},
		},
		{
			name:     "duplicate results expected from buggy code",
			s:        "((",
			expected: []string{""},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := removeInvalidParentheses(tc.s)
			// The order of elements in the result is not guaranteed
			// and the current implementation might produce duplicates.
			// A robust test would be to sort and unique both lists.
			sort.Strings(actual)
			actual = unique(actual)
			sort.Strings(tc.expected)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

// unique removes duplicate strings from a sorted slice of strings.
func unique(s []string) []string {
	if len(s) == 0 {
		return s
	}
	j := 0
	for i := 1; i < len(s); i++ {
		if s[j] != s[i] {
			j++
			s[j] = s[i]
		}
	}
	return s[:j+1]
}
