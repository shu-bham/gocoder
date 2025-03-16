package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	leetcode "gocoder/leetcode/daily"
	"testing"
)

func TestMinCapability(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, 5, leetcode.MinCapability([]int{2, 3, 5, 9}, 2))
	})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, 2, leetcode.MinCapability([]int{2, 7, 9, 3, 1}, 2))
	})
}
