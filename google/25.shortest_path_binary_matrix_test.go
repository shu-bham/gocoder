package google_test

import (
	"gocoder/google"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestPathBinaryMatrix(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name:     "Example 1",
			grid:     [][]int{{0, 1}, {1, 0}},
			expected: 2,
		},
		{
			name:     "Example 2",
			grid:     [][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}},
			expected: 4,
		},
		{
			name:     "No Path",
			grid:     [][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 0}},
			expected: -1,
		},
		{
			name:     "Destination blocked",
			grid:     [][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 1}},
			expected: -1,
		},
		{
			name:     "1x1 grid success",
			grid:     [][]int{{0}},
			expected: 1,
		},
		{
			name:     "1x1 grid fail",
			grid:     [][]int{{1}},
			expected: -1,
		},
		{
			name:     "failing",
			grid:     [][]int{{0, 0, 0}, {0, 0, 0}, {0, 1, 0}},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := google.ShortestPathBinaryMatrix(tt.grid)
			assert.Equal(t, tt.expected, result)
		})
	}
}
