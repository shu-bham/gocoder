package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	leetcode "gocoder/leetcode"
	"testing"
)

func TestValidPath(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		validPath := leetcode.ValidPath(3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2)
		assert.Equal(t, true, validPath)
	})

	t.Run("t2", func(t *testing.T) {
		validPath := leetcode.ValidPath(6, [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, 0, 5)
		assert.Equal(t, false, validPath)
	})
}
