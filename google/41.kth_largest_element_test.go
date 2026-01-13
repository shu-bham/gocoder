package google

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "test case 1",
			nums: []int{3, 2, 1, 5, 6, 4},
			k:    2,
			want: 5,
		},
		{
			name: "test case 2",
			nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:    4,
			want: 4,
		},
		{
			name: "single element",
			nums: []int{1},
			k:    1,
			want: 1,
		},
		{
			name: "all elements same",
			nums: []int{7, 7, 7, 7},
			k:    3,
			want: 7,
		},
		{
			name: "k is 1",
			nums: []int{99, 1, 2, 3},
			k:    1,
			want: 99,
		},
		{
			name: "k is length of array",
			nums: []int{5, 2, 8, 1, 9},
			k:    5,
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindKthLargest(tt.nums, tt.k)
			assert.Equal(t, tt.want, got)
		})
	}
}
