package google_test

import (
	"gocoder/google"
	"testing"
)

func TestLengthOfLIS(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{10, 9, 2, 5, 3, 7, 101, 18},
			expected: 4,
		},
		{
			name:     "Example 2",
			nums:     []int{0, 1, 0, 3, 2, 3},
			expected: 4,
		},
		{
			name:     "Example 3",
			nums:     []int{7, 7, 7, 7, 7, 7, 7},
			expected: 1,
		},
		{
			name:     "Empty slice",
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "Single element",
			nums:     []int{5},
			expected: 1,
		},
		{
			name:     "Sorted slice",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 5,
		},
		{
			name:     "Reverse sorted slice",
			nums:     []int{5, 4, 3, 2, 1},
			expected: 1,
		},
		{
			name:     "Complex case",
			nums:     []int{3, 5, 6, 2, 5, 4, 19, 5, 6, 7, 12},
			expected: 6, // [3,5,6,7,12] or [3,5,6,19] or [3,4,5,6,7,12]
		},
		{
			name:     "custom1",
			nums:     []int{1, 3, 6, 7, 9, 4, 10, 5, 6},
			expected: 6,
		},
		{
			name:     "custom2",
			nums:     []int{1},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := google.LengthOfLIS(tc.nums); got != tc.expected {
				t.Errorf("LengthOfLIS(%v) = %d; want %d", tc.nums, got, tc.expected)
			}
		})
	}
}

func TestLengthOfLISv2(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{10, 9, 2, 5, 3, 7, 101, 18},
			expected: 4,
		},
		{
			name:     "Example 2",
			nums:     []int{0, 1, 0, 3, 2, 3},
			expected: 4,
		},
		{
			name:     "Example 3",
			nums:     []int{7, 7, 7, 7, 7, 7, 7},
			expected: 1,
		},
		{
			name:     "Empty slice",
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "Single element",
			nums:     []int{5},
			expected: 1,
		},
		{
			name:     "Sorted slice",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 5,
		},
		{
			name:     "Reverse sorted slice",
			nums:     []int{5, 4, 3, 2, 1},
			expected: 1,
		},
		{
			name:     "Complex case",
			nums:     []int{3, 5, 6, 2, 5, 4, 19, 5, 6, 7, 12},
			expected: 6, // [3,5,6,7,12] or [3,5,6,19] or [3,4,5,6,7,12]
		},
		{
			name:     "custom1",
			nums:     []int{1, 3, 6, 7, 9, 4, 10, 5, 6},
			expected: 6,
		},
		{
			name:     "custom2",
			nums:     []int{1},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := google.LengthOfLISv2(tc.nums); got != tc.expected {
				t.Errorf("LengthOfLIS(%v) = %d; want %d", tc.nums, got, tc.expected)
			}
		})
	}
}
