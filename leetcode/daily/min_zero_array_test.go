package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	leetcode "gocoder/leetcode/daily"
	"testing"
)

func TestMinZeroArray(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		k := leetcode.MinZeroArray([]int{2, 0, 2}, [][]int{{0, 2, 1}, {0, 2, 1}, {1, 1, 3}})
		assert.Equal(t, 2, k)
	})

	t.Run("t2", func(t *testing.T) {
		k := leetcode.MinZeroArray([]int{2, 7, 6, 8, 6, 1, 6, 5, 3, 1}, [][]int{{0, 4, 1}, {2, 7, 3}, {4, 9, 3}, {1, 6, 2}, {0, 5, 4}, {3, 7, 4}, {1, 3, 1}})
		assert.Equal(t, 5, k)
	})
}
