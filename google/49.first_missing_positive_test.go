package google

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstMissingPositive(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 0},
			expected: 3,
		},
		{
			name:     "Example 2",
			nums:     []int{3, 4, -1, 1},
			expected: 2,
		},
		{
			name:     "Example 3",
			nums:     []int{7, 8, 9, 11, 12},
			expected: 1,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: 1,
		},
		{
			name:     "All negatives",
			nums:     []int{-1, -2, -3},
			expected: 1,
		},
		{
			name:     "Contains zero",
			nums:     []int{0},
			expected: 1,
		},
		{
			name:     "Sequence with gap",
			nums:     []int{1, 2, 4, 5},
			expected: 3,
		},
		{
			name:     "Duplicates",
			nums:     []int{1, 1, 2, 0},
			expected: 3,
		},
		{
			name:     "Large numbers",
			nums:     []int{100, 200, 300},
			expected: 1,
		},
		{
			name:     "Mixed numbers",
			nums:     []int{1, 2, 3, -5, 0, 4},
			expected: 5,
		},
		{
			name:     "Many duplicates and a missing in between",
			nums:     []int{1, 1, 2, 2, 4, 4, 5},
			expected: 3,
		},
		{
			name:     "Missing is n+1",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 6,
		},
		{
			name:     "Single positive number",
			nums:     []int{1},
			expected: 2,
		},
		{
			name:     "Single negative number",
			nums:     []int{-5},
			expected: 1,
		},
		{
			name:     "Zeros and negatives, but 1 is missing",
			nums:     []int{0, -1, -2, 2, 3},
			expected: 1,
		},
		{
			name:     "Zero and one",
			nums:     []int{0, 1},
			expected: 2,
		},
		{
			name:     "Large array, many duplicates, missing in middle",
			nums:     []int{1, 2, 3, 1, 2, 3, 5, 6, 7, 8, 9, 10},
			expected: 4,
		},
		{
			name:     "Large array, all present",
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: 11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, firstMissingPositive(tt.nums))
		})
	}
}
