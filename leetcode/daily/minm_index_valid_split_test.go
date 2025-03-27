package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	leetcode "gocoder/leetcode/daily"
	"testing"
)

func TestMinimumIndexOfValidSplit(t *testing.T) {

	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, 4, leetcode.MinimumIndexOfValidSplit([]int{2, 1, 3, 1, 1, 1, 7, 1, 2, 1}))
	})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, 2, leetcode.MinimumIndexOfValidSplit([]int{1, 2, 2, 2}))
	})

	t.Run("t3", func(t *testing.T) {
		assert.Equal(t, -1, leetcode.MinimumIndexOfValidSplit([]int{3, 3, 3, 3, 7, 2, 2}))
	})
}

func TestMinimumIndexOfValidSplitV2(t *testing.T) {

	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, 4, leetcode.MinimumIndexOfValidSplitV2([]int{2, 1, 3, 1, 1, 1, 7, 1, 2, 1}))
	})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, 2, leetcode.MinimumIndexOfValidSplitV2([]int{1, 2, 2, 2}))
	})

	t.Run("t3", func(t *testing.T) {
		assert.Equal(t, -1, leetcode.MinimumIndexOfValidSplitV2([]int{3, 3, 3, 3, 7, 2, 2}))
	})
}
