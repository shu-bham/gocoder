package google

import (
	"reflect"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		expected  [][]int
	}{
		{
			name:      "empty intervals",
			intervals: [][]int{},
			expected:  [][]int{},
		},
		{
			name:      "single interval",
			intervals: [][]int{{1, 4}},
			expected:  [][]int{{1, 4}},
		},
		{
			name:      "non-overlapping intervals",
			intervals: [][]int{{1, 2}, {3, 4}},
			expected:  [][]int{{1, 2}, {3, 4}},
		},
		{
			name:      "overlapping intervals",
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			expected:  [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:      "intervals fully containing others",
			intervals: [][]int{{1, 10}, {2, 5}},
			expected:  [][]int{{1, 10}},
		},
		{
			name:      "intervals merging into one large interval",
			intervals: [][]int{{1, 4}, {0, 4}},
			expected:  [][]int{{0, 4}},
		},
		{
			name:      "intervals needing sorting and then merging",
			intervals: [][]int{{4, 5}, {1, 4}, {0, 0}},
			expected:  [][]int{{0, 0}, {1, 5}},
		},
		{
			name:      "complex overlapping intervals",
			intervals: [][]int{{1, 4}, {0, 0}, {4, 5}},
			expected:  [][]int{{0, 0}, {1, 5}},
		},
		{
			name:      "another complex overlapping intervals",
			intervals: [][]int{{1, 4}, {0, 2}, {3, 5}},
			expected:  [][]int{{0, 5}},
		},
		{
			name:      "intervals with negative numbers",
			intervals: [][]int{{-10, -5}, {-6, -1}, {0, 1}},
			expected:  [][]int{{-10, -1}, {0, 1}},
		},
		{
			name:      "intervals with same start and end",
			intervals: [][]int{{1, 4}, {4, 5}},
			expected:  [][]int{{1, 5}},
		},
		{
			name:      "intervals with same start",
			intervals: [][]int{{1, 4}, {1, 5}},
			expected:  [][]int{{1, 5}},
		},
		{
			name:      "intervals with same end",
			intervals: [][]int{{1, 5}, {2, 5}},
			expected:  [][]int{{1, 5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeIntervals(tt.intervals)
			// Sort both actual and expected for comparison, as the order of merged intervals might vary
			// but the content should be the same. This is assuming the MergeIntervals implementation
			// might not guarantee a specific order if multiple merges happen in parallel,
			// though typically it would be sorted. For a robust test, it's good to sort.
			// However, a correct merge function should always return sorted intervals.
			// For simplicity, we'll assume the implementation sorts them and check directly.
			// If the problem statement guarantees sorted output, direct comparison is fine.
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("MergeIntervals() got = %v, want %v", got, tt.expected)
			}
		})
	}
}
