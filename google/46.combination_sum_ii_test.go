package google

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	testCases := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{
		{
			name:       "example 1",
			candidates: []int{10, 1, 2, 7, 6, 1, 5},
			target:     8,
			want: [][]int{
				{1, 1, 6},
				{1, 2, 5},
				{1, 7},
				{2, 6},
			},
		},
		{
			name:       "example 2",
			candidates: []int{2, 5, 2, 1, 2},
			target:     5,
			want: [][]int{
				{1, 2, 2},
				{5},
			},
		},
		{
			name:       "no solution",
			candidates: []int{2, 3, 5},
			target:     1,
			want:       [][]int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := combinationSum2(tc.candidates, tc.target)
			for i := range got {
				sort.Ints(got[i])
			}
			for i := range tc.want {
				sort.Ints(tc.want[i])
			}
			assert.ElementsMatch(t, tc.want, got)
		})
	}
}
