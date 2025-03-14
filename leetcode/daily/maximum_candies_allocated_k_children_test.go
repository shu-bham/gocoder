package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	leetcode "gocoder/leetcode/daily"
	"testing"
)

func TestMaximumCandies(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, 5, leetcode.MaximumCandies([]int{3, 4, 5, 10}, 3))
	})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, 7, leetcode.MaximumCandies([]int{3, 4, 4, 21}, 3))
	})

	t.Run("t3", func(t *testing.T) {
		assert.Equal(t, 0, leetcode.MaximumCandies([]int{3, 4, 4}, 12))
	})
}

func TestMaximumCandiesV2(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, 5, leetcode.MaximumCandiesV2([]int{3, 4, 5, 10}, 3))
	})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, 7, leetcode.MaximumCandiesV2([]int{3, 4, 4, 21}, 3))
	})

	t.Run("t3", func(t *testing.T) {
		assert.Equal(t, 0, leetcode.MaximumCandiesV2([]int{3, 4, 4}, 12))
	})

	t.Run("t4", func(t *testing.T) {
		assert.Equal(t, 5, leetcode.MaximumCandiesV2([]int{5, 8, 6}, 3))
	})
}
