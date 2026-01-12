package google

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePathsWithObstacles(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "Example 1",
			grid: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
			expected: 2,
		},
		{
			name: "Example 2",
			grid: [][]int{
				{0, 1},
				{0, 0},
			},
			expected: 1,
		},
		{
			name: "Obstacle at start",
			grid: [][]int{
				{1, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			expected: 0,
		},
		{
			name: "Obstacle at end",
			grid: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 1},
			},
			expected: 0,
		},
		{
			name: "No path",
			grid: [][]int{
				{0, 1},
				{1, 0},
			},
			expected: 0,
		},
		{
			name:     "Single cell no obstacle",
			grid:     [][]int{{0}},
			expected: 1,
		},
		{
			name:     "Single cell with obstacle",
			grid:     [][]int{{1}},
			expected: 0,
		},
		{
			name: "Larger grid",
			grid: [][]int{
				{0, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 0},
			},
			expected: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := UniquePathsWithObstacles(tt.grid)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
