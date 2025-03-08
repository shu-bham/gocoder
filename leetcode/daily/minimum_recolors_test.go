package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	leetcode "gocoder/leetcode/daily"
	"testing"
)

func TestMinimumRecolors(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, 3, leetcode.MinimumRecolors("WBBWWBBWBW", 7))
	})

	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, 0, leetcode.MinimumRecolors("WBWBBBW", 2))
	})
}
