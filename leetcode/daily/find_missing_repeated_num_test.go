package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	leetcode "gocoder/leetcode/daily"
	"testing"
)

func TestFindMissingAndRepeatedValues(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, []int{2, 4}, leetcode.FindMissingAndRepeatedValues([][]int{{1, 3}, {2, 2}}))
	})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, []int{9, 5}, leetcode.FindMissingAndRepeatedValues([][]int{{9, 1, 7}, {8, 9, 2}, {3, 4, 6}}))
	})
}

func TestFindMissingAndRepeatedValuesV2(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, []int{2, 4}, leetcode.FindMissingAndRepeatedValuesV2([][]int{{1, 3}, {2, 2}}))
	})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, []int{9, 5}, leetcode.FindMissingAndRepeatedValuesV2([][]int{{9, 1, 7}, {8, 9, 2}, {3, 4, 6}}))
	})
}
