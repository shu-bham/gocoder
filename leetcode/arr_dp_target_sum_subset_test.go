package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/leetcode"
	"testing"
)

func TestTargetSumSubset(t *testing.T) {
	arr, target := []int{4, 2, 7, 1, 3}, 12
	assert.Equal(t, true, leetcode.TargetSumSubset(arr, target))

	arr, target = []int{1, 0, 0, 3, 2, 5}, 6
	assert.Equal(t, true, leetcode.TargetSumSubset(arr, target))

	arr, target = []int{0, 0, 3, 2, 5}, 4
	assert.Equal(t, false, leetcode.TargetSumSubset(arr, target))

	arr, target = []int{0, 2, 1}, 4
	assert.Equal(t, false, leetcode.TargetSumSubset(arr, target))
}
