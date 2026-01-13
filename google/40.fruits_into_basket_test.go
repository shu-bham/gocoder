package google

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTotalFruit(t *testing.T) {
	tests := []struct {
		name   string
		fruits []int
		want   int
	}{
		{
			name:   "test case 1",
			fruits: []int{1, 2, 1},
			want:   3,
		},
		{
			name:   "test case 2",
			fruits: []int{0, 1, 2, 2},
			want:   3,
		},
		{
			name:   "test case 3",
			fruits: []int{1, 2, 3, 2, 2},
			want:   4,
		},
		{
			name:   "test case 4",
			fruits: []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4},
			want:   5,
		},
		{
			name:   "single fruit",
			fruits: []int{1},
			want:   1,
		},
		{
			name:   "all same fruits",
			fruits: []int{1, 1, 1, 1, 1},
			want:   5,
		},
		{
			name:   "empty input",
			fruits: []int{},
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TotalFruit(tt.fruits)
			assert.Equal(t, tt.want, got)
		})
	}
}
