package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	leetcode "gocoder/leetcode/daily"
	"testing"
)

func TestMaximumCountPositiveOrNegative(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, 3, leetcode.MaximumCountPositiveOrNegative([]int{-2, -1, -1, 1, 2, 3}))
	})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, 3, leetcode.MaximumCountPositiveOrNegative([]int{-3, -2, -1, 0, 0, 1, 2}))
	})

	t.Run("t3", func(t *testing.T) {
		assert.Equal(t, 4, leetcode.MaximumCountPositiveOrNegative([]int{5, 20, 66, 1314}))
	})

	t.Run("t4", func(t *testing.T) {
		assert.Equal(t, 0, leetcode.MaximumCountPositiveOrNegative([]int{0, 0, 0, 0, 0, 0}))
	})

	t.Run("t5", func(t *testing.T) {
		assert.Equal(t, 2, leetcode.MaximumCountPositiveOrNegative([]int{0, 0, 0, 0, 1, 2}))
	})

	t.Run("t6", func(t *testing.T) {
		assert.Equal(t, 1, leetcode.MaximumCountPositiveOrNegative([]int{-1, 0, 0, 0, 0, 0}))
	})
}
