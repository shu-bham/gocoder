package google_test

import (
	"gocoder/google"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected string
	}{
		{
			name:     "Example 1",
			s:        "babad",
			expected: "bab",
		},
		{
			name:     "Example 2",
			s:        "cbbd",
			expected: "bb",
		},
		{
			name:     "Single character",
			s:        "a",
			expected: "a",
		},
		{
			name:     "Empty string",
			s:        "",
			expected: "",
		},
		{
			name:     "Two same characters",
			s:        "aa",
			expected: "aa",
		},
		{
			name:     "Palindrome string",
			s:        "racecar",
			expected: "racecar",
		},
		{
			name:     "No palindrome",
			s:        "abcdef",
			expected: "f",
		},
		{
			name:     "Long string",
			s:        "bananas",
			expected: "anana",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := google.LongestPalindrome(tc.s); got != tc.expected {
				if tc.s == "babad" && (got == "bab" || got == "aba") {
					return
				}
				t.Errorf("LongestPalindrome(%q) = %q; want %q", tc.s, got, tc.expected)
			}
		})
	}
}
