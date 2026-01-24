package google

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToRoman(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected string
	}{
		{
			name:     "Test Case 1: 3",
			input:    3,
			expected: "III",
		},
		{
			name:     "Test Case 2: 58",
			input:    58,
			expected: "LVIII",
		},
		{
			name:     "Test Case 3: 1994",
			input:    1994,
			expected: "MCMXCIV",
		},
		{
			name:     "Test Case 4: 9",
			input:    9,
			expected: "IX",
		},
		{
			name:     "Test Case 5: 4",
			input:    4,
			expected: "IV",
		},
		{
			name:     "Test Case 6: 1",
			input:    1,
			expected: "I",
		},
		{
			name:     "Test Case 7: 3999",
			input:    3999,
			expected: "MMMCMXCIX",
		},
		{
			name:     "Test Case 8: 0",
			input:    0,
			expected: "",
		},
		{
			name:     "Test Case 9: 4000 (invalid)",
			input:    4000,
			expected: "",
		},
		{
			name:     "Test Case 10: -1 (invalid)",
			input:    -1,
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := intToRoman(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}
