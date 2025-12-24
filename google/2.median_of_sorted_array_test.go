package google

import (
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		nums1    []int
		nums2    []int
		expected float64
	}{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{0, 0}, []int{0, 0}, 0.0},
		{[]int{}, []int{1}, 1.0},
		{[]int{2}, []int{}, 2.0},
		{[]int{3}, []int{-2, -1}, -1.0},
	}

	for _, test := range tests {
		result := FindMedianSortedArrays(test.nums1, test.nums2)
		if result != test.expected {
			t.Errorf("For nums1: %v, nums2: %v, expected: %f, got: %f", test.nums1, test.nums2, test.expected, result)
		}
	}
}

func TestFindMedianSortedArraysV2(t *testing.T) {
	tests := []struct {
		nums1    []int
		nums2    []int
		expected float64
	}{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{0, 0}, []int{0, 0}, 0.0},
		{[]int{}, []int{1}, 1.0},
		{[]int{2}, []int{}, 2.0},
		{[]int{3}, []int{-2, -1}, -1.0},
	}

	for _, test := range tests {
		result := FindMedianSortedArraysV2(test.nums1, test.nums2)
		if result != test.expected {
			t.Errorf("For nums1: %v, nums2: %v, expected: %f, got: %f", test.nums1, test.nums2, test.expected, result)
		}
	}
}

func TestFindMedianSortedArraysV3(t *testing.T) {
	tests := []struct {
		nums1    []int
		nums2    []int
		expected float64
	}{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{0, 0}, []int{0, 0}, 0.0},
		{[]int{}, []int{1}, 1.0},
		{[]int{2}, []int{}, 2.0},
		{[]int{3}, []int{-2, -1}, -1.0},
	}

	for _, test := range tests {
		result := FindMedianSortedArraysV3(test.nums1, test.nums2)
		if result != test.expected {
			t.Errorf("For nums1: %v, nums2: %v, expected: %f, got: %f", test.nums1, test.nums2, test.expected, result)
		}
	}
}
