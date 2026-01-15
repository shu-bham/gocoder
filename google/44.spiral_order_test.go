package google

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpiralOrder(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   []int
	}{
		{
			name:   "Empty Matrix",
			matrix: [][]int{},
			want:   []int{},
		},
		{
			name:   "Single Element Matrix",
			matrix: [][]int{{1}},
			want:   []int{1},
		},
		{
			name:   "2x2 Matrix",
			matrix: [][]int{{1, 2}, {4, 3}},
			want:   []int{1, 2, 3, 4},
		},
		{
			name:   "3x3 Matrix",
			matrix: [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}},
			want:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:   "3x4 Matrix",
			matrix: [][]int{{1, 2, 3, 4}, {10, 11, 12, 5}, {9, 8, 7, 6}},
			want:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		},
		{
			name:   "1x5 Matrix",
			matrix: [][]int{{1, 2, 3, 4, 5}},
			want:   []int{1, 2, 3, 4, 5},
		},
		{
			name:   "5x1 Matrix",
			matrix: [][]int{{1}, {2}, {3}, {4}, {5}},
			want:   []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SpiralOrder(tt.matrix)
			assert.Equal(t, tt.want, got)
		})
	}
}
