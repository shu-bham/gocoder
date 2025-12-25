package google_test

import (
	"gocoder/google"
	"testing"
)

func TestMaxArea(t *testing.T) {
	testCases := []struct {
		name     string
		height   []int
		expected int
	}{
		{
			name:     "Example 1",
			height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			name:     "Example 2",
			height:   []int{1, 1},
			expected: 1,
		},
		{
			name:     "Empty input",
			height:   []int{},
			expected: 0,
		},
		{
			name:     "Two elements",
			height:   []int{4, 3},
			expected: 3,
		},
		{
			name:     "Decreasing heights",
			height:   []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: 20,
		},
		{
			name:     "Increasing heights",
			height:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: 20,
		},
		{
			name:     "Peaks at ends",
			height:   []int{2, 3, 4, 5, 18, 17, 6},
			expected: 17,
		},
		{
			name:     "All same heights",
			height:   []int{5, 5, 5, 5, 5},
			expected: 20,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := google.MaxArea(tc.height); got != tc.expected {
				t.Errorf("MaxArea(%v) = %d; want %d", tc.height, got, tc.expected)
			}
		})
	}
}

func TestMaxAreaV2(t *testing.T) {
	testCases := []struct {
		name     string
		height   []int
		expected int
	}{
		{
			name:     "Example 1",
			height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			name:     "Example 2",
			height:   []int{1, 1},
			expected: 1,
		},
		{
			name:     "Empty input",
			height:   []int{},
			expected: 0,
		},
		{
			name:     "Two elements",
			height:   []int{4, 3},
			expected: 3,
		},
		{
			name:     "Decreasing heights",
			height:   []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: 20,
		},
		{
			name:     "Increasing heights",
			height:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: 20,
		},
		{
			name:     "Peaks at ends",
			height:   []int{2, 3, 4, 5, 18, 17, 6},
			expected: 17,
		},
		{
			name:     "All same heights",
			height:   []int{5, 5, 5, 5, 5},
			expected: 20,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := google.MaxAreaV2(tc.height); got != tc.expected {
				t.Errorf("MaxAreaV2(%v) = %d; want %d", tc.height, got, tc.expected)
			}
		})
	}
}
