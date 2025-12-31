package google

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCanSeePersonsCount(t *testing.T) {
	testCases := []struct {
		heights []int
		want    []int
	}{
		{[]int{10, 6, 8, 5, 11, 9}, []int{3, 1, 2, 1, 1, 0}},
		{[]int{5, 1, 2, 3, 10}, []int{4, 1, 1, 1, 0}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("heights: %v", tc.heights), func(t *testing.T) {
			got := CanSeePersonsCount(tc.heights)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
