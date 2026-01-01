package google

import (
	"reflect"
	"testing"
)

func TestStrStr(t *testing.T) {
	tests := []struct {
		name     string
		haystack string
		needle   string
		expected int
	}{
		{
			name:     "Example 1: Needle at beginning",
			haystack: "sadbutsad",
			needle:   "sad",
			expected: 0,
		},
		{
			name:     "Example 2: Needle not found",
			haystack: "leetcode",
			needle:   "leeto",
			expected: -1,
		},
		{
			name:     "Example 3: Needle in middle",
			haystack: "hello",
			needle:   "ll",
			expected: 2,
		},
		{
			name:     "Case 4: Needle not found (bba in aaaaa)",
			haystack: "aaaaa",
			needle:   "bba",
			expected: -1,
		},
		{
			name:     "Case 6: Empty haystack, non-empty needle",
			haystack: "",
			needle:   "a",
			expected: -1,
		},
		{
			name:     "Case 7: Non-empty haystack, empty needle",
			haystack: "a",
			needle:   "",
			expected: 0,
		},
		{
			name:     "Case 8: Needle longer than haystack",
			haystack: "abc",
			needle:   "abcd",
			expected: -1,
		},
		{
			name:     "Case 9: Complex match (mississippi, issip)",
			haystack: "mississippi",
			needle:   "issip",
			expected: 4,
		},
		{
			name:     "Case 10: Complex match (mississippi, pi)",
			haystack: "mississippi",
			needle:   "pi",
			expected: 9,
		},
		{
			name:     "Case 11: Needle at beginning of long repeating string",
			haystack: "aaaabaaa",
			needle:   "aaaa",
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := StrStr(test.haystack, test.needle)
			if result != test.expected {
				t.Errorf("For haystack: %s, needle: %s, expected %d, but got %d", test.haystack, test.needle, test.expected, result)
			}
		})
	}
}

func TestKMPSearch(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		pattern  string
		expected []int
	}{
		{
			name:     "No matches",
			text:     "abcdefg",
			pattern:  "xyz",
			expected: []int{},
		},
		{
			name:     "One match",
			text:     "hello world",
			pattern:  "world",
			expected: []int{6},
		},
		{
			name:     "Multiple matches",
			text:     "ababa",
			pattern:  "aba",
			expected: []int{0, 2},
		},
		{
			name:     "Overlapping matches",
			text:     "ababab",
			pattern:  "abab",
			expected: []int{0, 2},
		},
		{
			name:     "Empty pattern",
			text:     "abcdef",
			pattern:  "",
			expected: []int{},
		},
		{
			name:     "Empty text",
			text:     "",
			pattern:  "abc",
			expected: []int{},
		},
		{
			name:     "Pattern longer than text",
			text:     "abc",
			pattern:  "abcdef",
			expected: []int{},
		},
		{
			name:     "Complex pattern with repetition",
			text:     "AABAACAADAABAABA",
			pattern:  "AABA",
			expected: []int{0, 9, 12},
		},
		{
			name:     "Pattern at end",
			text:     "find the needle at the end",
			pattern:  "end",
			expected: []int{24},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := KMPSearch(test.text, test.pattern)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("For text: '%s', pattern: '%s', expected %v, but got %v", test.text, test.pattern, test.expected, result)
			}
		})
	}
}
