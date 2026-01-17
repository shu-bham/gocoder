package google

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindXSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		x    int
		want []int64
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 3, 4, 5},
			k:    3,
			x:    2,
			want: []int64{5, 7, 9}, // Subarrays: [1,2,3] -> top 2 sum = 2+3=5, [2,3,4] -> top 2 sum = 3+4=7, [3,4,5] -> top 2 sum = 4+5=9
		},
		{
			name: "Example 2",
			nums: []int{10, 8, 10, 9},
			k:    2,
			x:    1,
			want: []int64{10, 10, 10}, // Subarrays: [10,8] -> top 1 sum = 10, [8,10] -> top 1 sum = 10, [10,9] -> top 1 sum = 10
		},
		{
			name: "Small k",
			nums: []int{1, 2, 3},
			k:    1,
			x:    1,
			want: []int64{1, 2, 3},
		},
		{
			name: "x greater than unique elements in window",
			nums: []int{1, 1, 2, 2, 3},
			k:    3,
			x:    5,                // x is larger than 3 unique elements in any window
			want: []int64{4, 5, 7}, // [1,1,2] -> 1*2+2=4; [1,2,2] -> 1+2*2=5; [2,2,3] -> 2*2+3=7
		},
		{
			name: "Duplicate elements, x=1",
			nums: []int{1, 5, 2, 5, 3},
			k:    3,
			x:    1,
			want: []int64{5, 10, 5}, // [1,5,2]->5; [5,2,5]->5; [2,5,3]->5
		},
		{
			name: "Single element window",
			nums: []int{10},
			k:    1,
			x:    1,
			want: []int64{10},
		},
		{
			name: "Multiple occurrences of same value",
			nums: []int{1, 1, 1, 1, 1},
			k:    3,
			x:    1,
			want: []int64{3, 3, 3},
		},
		{
			name: "All same values, large x",
			nums: []int{7, 7, 7, 7, 7},
			k:    3,
			x:    5,
			want: []int64{21, 21, 21}, // Window [7,7,7], 7*3 = 21
		},
		{
			name: "Example",
			nums: []int{1, 2, 3, 4, 5, 6},
			k:    6,
			x:    1,
			want: []int64{6},
		},
		{
			name: "Example2",
			nums: []int{1, 1, 2, 2, 3, 4, 2, 3},
			k:    6,
			x:    2,
			want: []int64{6, 10, 12},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findXSum(tt.nums, tt.k, tt.x)
			assert.Equal(t, tt.want, got)
		})
	}
}
