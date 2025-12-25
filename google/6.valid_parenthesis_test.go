package google_test

import (
	"gocoder/google"
	"testing"
)

func TestIsValid(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected bool
	}{
		{
			name:     "Valid parentheses 1",
			s:        "()",
			expected: true,
		},
		{
			name:     "Valid parentheses 2",
			s:        "()[]{}",
			expected: true,
		},
		{
			name:     "Invalid parentheses 1",
			s:        "(]",
			expected: false,
		},
		{
			name:     "Invalid parentheses 2",
			s:        "([)]",
			expected: false,
		},
		{
			name:     "Valid parentheses 3",
			s:        "{[]}",
			expected: true,
		},
		{
			name:     "Empty string",
			s:        "",
			expected: true,
		},
		{
			name:     "Open bracket only",
			s:        "(",
			expected: false,
		},
		{
			name:     "Close bracket only",
			s:        ")",
			expected: false,
		},
		{
			name:     "Mismatched bracket",
			s:        "{[",
			expected: false,
		},
		{
			name:     "Mismatched bracket 2",
			s:        "]",
			expected: false,
		},
		{
			name:     "Mismatched bracket 3",
			s:        "((",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := google.IsValid(tc.s); got != tc.expected {
				t.Errorf("IsValid(%q) = %v; want %v", tc.s, got, tc.expected)
			}
		})
	}
}
