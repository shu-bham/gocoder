package google

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		expected   [][]int
	}{
		{
			name:       "example 1",
			candidates: []int{2, 3, 6, 7},
			target:     7,
			expected:   [][]int{{2, 2, 3}, {7}},
		},
		{
			name:       "example 2",
			candidates: []int{2, 3, 5},
			target:     8,
			expected:   [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name:       "no solution",
			candidates: []int{2},
			target:     1,
			expected:   [][]int{},
		},
		{
			name:       "single candidate",
			candidates: []int{1},
			target:     3,
			expected:   [][]int{{1, 1, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := combinationSum(tt.candidates, tt.target)
			// Sort the inner slices and the outer slice for consistent comparison
			for i := range actual {
				sort.Ints(actual[i])
			}
			sort.Slice(actual, func(i, j int) bool {
				for x := 0; x < len(actual[i]) && x < len(actual[j]); x++ {
					if actual[i][x] != actual[j][x] {
						return actual[i][x] < actual[j][x]
					}
				}
				return len(actual[i]) < len(actual[j])
			})

			for i := range tt.expected {
				sort.Ints(tt.expected[i])
			}
			sort.Slice(tt.expected, func(i, j int) bool {
				for x := 0; x < len(tt.expected[i]) && x < len(tt.expected[j]); x++ {
					if tt.expected[i][x] != tt.expected[j][x] {
						return tt.expected[i][x] < tt.expected[j][x]
					}
				}
				return len(tt.expected[i]) < len(tt.expected[j])
			})

			if len(actual) == 0 && len(tt.expected) == 0 {
				return // Both are empty, which is a pass
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}
