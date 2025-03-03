package util_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/util"
	"testing"
)

func TestQuickSort(t *testing.T) {
	// Test case 1: Regular unsorted array
	arr1 := []int{34, 7, 23, 32, 5, 62}
	expected1 := []int{5, 7, 23, 32, 34, 62}
	util.QuickSort(arr1)
	assert.Equal(t, expected1, arr1, "Array should be sorted")

	// Test case 2: Array with negative numbers
	arr2 := []int{-3, 10, -1, 5, 0}
	expected2 := []int{-3, -1, 0, 5, 10}
	util.QuickSort(arr2)
	assert.Equal(t, expected2, arr2, "Array with negatives should be sorted")

	// Test case 3: Array with duplicate numbers
	arr3 := []int{4, 2, 2, 8, 3, 3, 1}
	expected3 := []int{1, 2, 2, 3, 3, 4, 8}
	util.QuickSort(arr3)
	assert.Equal(t, expected3, arr3, "Array with duplicates should be sorted")

	// Test case 4: Already sorted array
	arr4 := []int{1, 2, 3, 4, 5}
	expected4 := []int{1, 2, 3, 4, 5}
	util.QuickSort(arr4)
	assert.Equal(t, expected4, arr4, "Already sorted array should remain unchanged")

	// Test case 5: Array with a single element
	arr5 := []int{42}
	expected5 := []int{42}
	util.QuickSort(arr5)
	assert.Equal(t, expected5, arr5, "Single element array should remain unchanged")

	// Test case 6: Empty array
	arr6 := []int{}
	expected6 := []int{}
	util.QuickSort(arr6)
	assert.Equal(t, expected6, arr6, "Empty array should remain unchanged")
}
