package google_test

import (
	"gocoder/google"
	"testing"
)

func TestCombinationSum4(t *testing.T) {
	tests := []struct {
		coins    []int
		target   int
		expected int
	}{
		{[]int{1, 2, 3}, 4, 7},
		{[]int{9}, 3, 0},
		{[]int{1, 2}, 3, 3},
		{[]int{3, 1, 2}, 4, 7},
		{[]int{1}, 0, 1}, // Target 0 should always have 1 combination (empty set)
		{[]int{2, 3, 5}, 10, 14},
	}

	for _, test := range tests {
		result := google.CombinationSum4(test.coins, test.target)
		if result != test.expected {
			t.Errorf("For coins %v and target %d, expected %d, but got %d", test.coins, test.target, test.expected, result)
		}
	}
}
