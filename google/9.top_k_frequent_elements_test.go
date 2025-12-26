package google_test

import (
	"gocoder/google"
	"reflect"
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 1, 1, 2, 2, 3},
			k:        2,
			expected: []int{1, 2},
		},
		{
			name:     "Example 2",
			nums:     []int{1},
			k:        1,
			expected: []int{1},
		},
		{
			name:     "Multiple frequencies",
			nums:     []int{1, 2, 3, 1, 2, 1, 3, 3, 3},
			k:        2,
			expected: []int{3, 1}, // or [1,3], order doesn't matter, but the implementation might produce one consistently
		},
		{
			name:     "All same frequency",
			nums:     []int{1, 2, 3, 4, 5},
			k:        3,
			expected: []int{1, 2, 3}, // any 3 unique elements
		},
		{
			name:     "Negative numbers",
			nums:     []int{-1, -1, -2, -2, -2, 0},
			k:        2,
			expected: []int{-2, -1},
		},
		{
			name:     "Empty slice",
			nums:     []int{},
			k:        0,
			expected: []int{},
		},
		{
			name:     "k equals length of nums",
			nums:     []int{1, 2, 3},
			k:        3,
			expected: []int{1, 2, 3}, // order can vary
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := google.TopKFrequent(tc.nums, tc.k)
			sort.Ints(got) // Sort to handle cases where order of elements with same frequency doesn't matter
			sort.Ints(tc.expected)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("TopKFrequent(%v, %d) = %v; want %v", tc.nums, tc.k, got, tc.expected)
			}
		})
	}
}
