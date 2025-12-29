package google

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "target found multiple times",
			nums:     []int{5, 7, 7, 8, 8, 10},
			target:   8,
			expected: []int{3, 4},
		},
		{
			name:     "target found once",
			nums:     []int{5, 7, 7, 8, 8, 10},
			target:   7,
			expected: []int{1, 2},
		},
		{
			name:     "target not found",
			nums:     []int{5, 7, 7, 8, 8, 10},
			target:   6,
			expected: []int{-1, -1},
		},
		{
			name:     "empty array",
			nums:     []int{},
			target:   0,
			expected: []int{-1, -1},
		},
		{
			name:     "single element array, target found",
			nums:     []int{1},
			target:   1,
			expected: []int{0, 0},
		},
		{
			name:     "single element array, target not found",
			nums:     []int{1},
			target:   0,
			expected: []int{-1, -1},
		},
		{
			name:     "target at the beginning",
			nums:     []int{1, 2, 3},
			target:   1,
			expected: []int{0, 0},
		},
		{
			name:     "target at the end",
			nums:     []int{1, 2, 3},
			target:   3,
			expected: []int{2, 2},
		},
		{
			name:     "all elements are target",
			nums:     []int{3, 3, 3, 3, 3},
			target:   3,
			expected: []int{0, 4},
		},
		{
			name:     "target is smallest element",
			nums:     []int{1, 2, 3, 4, 5},
			target:   1,
			expected: []int{0, 0},
		},
		{
			name:     "target is largest element",
			nums:     []int{1, 2, 3, 4, 5},
			target:   5,
			expected: []int{4, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SearchRange(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("SearchRange() got = %v, want %v", got, tt.expected)
			}
		})
	}
}
