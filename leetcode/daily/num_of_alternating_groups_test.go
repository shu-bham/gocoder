package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	leetcode "gocoder/leetcode/daily"
	"testing"
)

func TestNumberOfAlternatingGroups(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		num := leetcode.NumberOfAlternatingGroups([]int{0, 1, 0, 0, 1, 0, 1}, 6)
		assert.Equal(t, 2, num)
	})

	t.Run("t2", func(t *testing.T) {
		num := leetcode.NumberOfAlternatingGroups([]int{0, 1, 0, 1, 0}, 3)
		assert.Equal(t, 3, num)
	})
}
