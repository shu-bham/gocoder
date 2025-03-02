package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	leetcode "gocoder/leetcode/daily"
	"testing"
)

func TestMergeArraysSum2D(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		arr1 := [][]int{{1, 2}, {2, 3}, {4, 5}}
		arr2 := [][]int{{1, 4}, {3, 2}, {4, 1}}
		sum := leetcode.MergeArraysSum(arr1, arr2)
		assert.Equal(t, [][]int{{1, 6}, {2, 3}, {3, 2}, {4, 6}}, sum)
	})
}
