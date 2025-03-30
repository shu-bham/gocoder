package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	leetcode "gocoder/leetcode/daily"
	"testing"
)

func TestMaximumScoreAfterApplyingOps(t *testing.T) {

	t.Run("t1", func(t *testing.T) {
		nums := []int{8, 3, 9, 3, 8}
		k := 2
		assert.Equal(t, 81, leetcode.MaximumScoreAfterApplyingOps(nums, k))
	})

	t.Run("t2", func(t *testing.T) {
		nums := []int{19, 12, 14, 6, 10, 18}
		k := 3
		assert.Equal(t, 4788, leetcode.MaximumScoreAfterApplyingOps(nums, k))
	})

	t.Run("t3", func(t *testing.T) {
		nums := []int{8, 3, 9, 3, 8}
		k := 2
		assert.Equal(t, 81, leetcode.MaximumScoreAfterApplyingOps(nums, k))
	})

	t.Run("t4", func(t *testing.T) {
		nums := []int{10, 13, 20, 25, 21, 24}
		k := 4
		assert.Equal(t, 264600, leetcode.MaximumScoreAfterApplyingOps(nums, k))
	})

	t.Run("t5", func(t *testing.T) {
		nums := []int{27, 30, 15, 22, 31}
		k := 6
		assert.Equal(t, 753300000, leetcode.MaximumScoreAfterApplyingOps(nums, k))
	})

	t.Run("t6", func(t *testing.T) {
		nums := []int{27, 30, 15, 22, 31}
		k := 7
		assert.Equal(t, 598999846, leetcode.MaximumScoreAfterApplyingOps(nums, k))
	})

	t.Run("t7", func(t *testing.T) {
		nums := []int{27, 30, 15, 22, 31}
		k := 8
		assert.Equal(t, 969995261, leetcode.MaximumScoreAfterApplyingOps(nums, k))
	})

	t.Run("t8", func(t *testing.T) {
		nums := []int{27, 30, 15, 22, 31}
		k := 9
		assert.Equal(t, 99857627, leetcode.MaximumScoreAfterApplyingOps(nums, k))

	})

	t.Run("t9", func(t *testing.T) {
		nums := []int{1, 1, 2, 18, 1, 9, 3, 1}
		k := 32
		assert.Equal(t, 346264255, leetcode.MaximumScoreAfterApplyingOps(nums, k))

	})

}
