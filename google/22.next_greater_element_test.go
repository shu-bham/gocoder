package google

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNextGreaterElement(t *testing.T) {
	testCases := []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{[]int{4, 1, 2}, []int{1, 3, 4, 2}, []int{-1, 3, -1}},
		{[]int{2, 4}, []int{1, 2, 3, 4}, []int{3, -1}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("nums1: %v, nums2: %v", tc.nums1, tc.nums2), func(t *testing.T) {
			got := nextGreaterElement(tc.nums1, tc.nums2)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
