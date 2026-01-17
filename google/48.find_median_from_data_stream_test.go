package google

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMedianFinder(t *testing.T) {
	tests := []struct {
		name     string
		actions  []string
		values   []int
		expected []float64
	}{
		{
			name:     "Example 1",
			actions:  []string{"addNum", "addNum", "findMedian", "addNum", "findMedian"},
			values:   []int{1, 2, 0, 3, 0}, // 0 is a placeholder for findMedian calls
			expected: []float64{-1.0, -1.0, 1.5, -1.0, 2.0},
		},
		{
			name:     "Single Element",
			actions:  []string{"addNum", "findMedian"},
			values:   []int{1, 0},
			expected: []float64{-1.0, 1.0},
		},
		{
			name:     "Odd Number of Elements",
			actions:  []string{"addNum", "addNum", "addNum", "findMedian"},
			values:   []int{1, 2, 3, 0},
			expected: []float64{-1.0, -1.0, -1.0, 2.0},
		},
		{
			name:     "Even Number of Elements",
			actions:  []string{"addNum", "addNum", "addNum", "addNum", "findMedian"},
			values:   []int{1, 2, 3, 4, 0},
			expected: []float64{-1.0, -1.0, -1.0, -1.0, 2.5},
		},
		{
			name:     "Negative Numbers",
			actions:  []string{"addNum", "addNum", "findMedian", "addNum", "findMedian"},
			values:   []int{-1, -2, 0, -3, 0},
			expected: []float64{-1.0, -1.0, -1.5, -1.0, -2.0},
		},
		{
			name:     "Mixed Numbers",
			actions:  []string{"addNum", "addNum", "addNum", "addNum", "addNum", "findMedian"},
			values:   []int{-1, 5, 2, 10, -3, 0},
			expected: []float64{-1.0, -1.0, -1.0, -1.0, -1.0, 2.0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mf := Constructor()
			var results []float64
			for i, action := range tt.actions {
				switch action {
				case "addNum":
					mf.AddNum(tt.values[i])
					results = append(results, -1.0) // Placeholder for addNum, not expecting a return value
				case "findMedian":
					results = append(results, mf.FindMedian())
				}
			}
			assert.InDeltaSlice(t, tt.expected, results, 0.001)
		})
	}
}
