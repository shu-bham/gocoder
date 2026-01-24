package google

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompress(t *testing.T) {
	testCases := []struct {
		name     string
		input    []byte
		expected int
		output   []byte
	}{
		{
			name:     "Example 1",
			input:    []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
			expected: 6,
			output:   []byte{'a', '2', 'b', '2', 'c', '3'},
		},
		{
			name:     "Example 2",
			input:    []byte{'a'},
			expected: 1,
			output:   []byte{'a'},
		},
		{
			name:     "Example 3",
			input:    []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
			expected: 4,
			output:   []byte{'a', 'b', '1', '2'},
		},
		{
			name:     "All same characters",
			input:    []byte{'a', 'a', 'a', 'a', 'a', 'a'},
			expected: 2,
			output:   []byte{'a', '6'},
		},
		{
			name:     "No compression",
			input:    []byte{'a', 'b', 'c', 'd', 'e'},
			expected: 5,
			output:   []byte{'a', 'b', 'c', 'd', 'e'},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Make a copy of the input slice to avoid modifying the original test case data.
			inputCopy := make([]byte, len(tc.input))
			copy(inputCopy, tc.input)

			result := compress(inputCopy)

			assert.Equal(t, tc.expected, result, "The returned length should match the expected length.")
			assert.Equal(t, tc.output, inputCopy[:result], "The compressed array content should match the expected output.")
		})
	}
}
