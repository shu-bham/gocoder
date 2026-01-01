package google

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		val      int
		expected []int
		k        int // expected length of the modified array
	}{
		{
			name:     "Example 1",
			nums:     []int{3, 2, 2, 3},
			val:      3,
			expected: []int{2, 2},
			k:        2,
		},
		{
			name:     "Example 2",
			nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:      2,
			expected: []int{0, 1, 3, 0, 4},
			k:        5,
		},
		{
			name:     "No elements to remove",
			nums:     []int{1, 2, 3, 4, 5},
			val:      6,
			expected: []int{1, 2, 3, 4, 5},
			k:        5,
		},
		{
			name:     "All elements to remove",
			nums:     []int{1, 1, 1, 1},
			val:      1,
			expected: []int{},
			k:        0,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			val:      0,
			expected: []int{},
			k:        0,
		},
		{
			name:     "Single element, remove it",
			nums:     []int{1},
			val:      1,
			expected: []int{},
			k:        0,
		},
		{
			name:     "Single element, don't remove it",
			nums:     []int{1},
			val:      2,
			expected: []int{1},
			k:        1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy of the original slice for modification, as RemoveElement modifies in-place
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)

			k := RemoveElement(numsCopy, tt.val)

			if k != tt.k {
				t.Errorf("For input nums: %v, val: %d, expected k: %d, got k: %d", tt.nums, tt.val, tt.k, k)
			}

			// Check if the elements in the modified part of the array match the expected elements
			for i := 0; i < k; i++ {
				if numsCopy[i] != tt.expected[i] {
					t.Errorf("For input nums: %v, val: %d, at index %d, expected: %d, got: %d", tt.nums, tt.val, i, tt.expected[i], numsCopy[i])
					break
				}
			}
		})
	}
}
