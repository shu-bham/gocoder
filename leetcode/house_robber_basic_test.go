package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/leetcode"
	"testing"
)

func TestRob(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, 4, leetcode.Rob([]int{1, 2, 3, 1}))
	})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, 12, leetcode.Rob([]int{2, 7, 9, 3, 1}))
	})
}
