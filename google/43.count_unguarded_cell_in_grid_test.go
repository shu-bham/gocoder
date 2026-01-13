package google

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountUnguarded(t *testing.T) {
	tests := []struct {
		name     string
		m        int
		n        int
		guards   [][]int
		walls    [][]int
		expected int
	}{
		{
			name:     "Example 1",
			m:        4,
			n:        6,
			guards:   [][]int{{0, 0}, {1, 1}, {2, 3}},
			walls:    [][]int{{0, 1}, {2, 2}, {1, 4}},
			expected: 7,
		},
		{
			name:     "Example 2",
			m:        3,
			n:        3,
			guards:   [][]int{{1, 1}},
			walls:    [][]int{},
			expected: 4,
		},
		{
			name:     "Single cell grid with a guard",
			m:        1,
			n:        1,
			guards:   [][]int{{0, 0}},
			walls:    [][]int{},
			expected: 0,
		},
		{
			name:     "Empty grid",
			m:        0,
			n:        0,
			guards:   [][]int{},
			walls:    [][]int{},
			expected: 0,
		},
		{
			name:     "No guards",
			m:        2,
			n:        2,
			guards:   [][]int{},
			walls:    [][]int{},
			expected: 4,
		},
		{
			name:     "All walls",
			m:        2,
			n:        2,
			guards:   [][]int{{0, 0}},
			walls:    [][]int{{0, 1}, {1, 0}, {1, 1}},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := CountUnguarded(tt.m, tt.n, tt.guards, tt.walls)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
