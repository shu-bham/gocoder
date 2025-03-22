package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	leetcode "gocoder/leetcode/daily"
	"testing"
)

func TestCountCompleteComponents(t *testing.T) {

	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, 1, leetcode.CountCompleteComponents(2, [][]int{{0, 1}}))
	})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, 2, leetcode.CountCompleteComponents(5, [][]int{{0, 1}, {2, 3}, {3, 4}, {4, 2}}))
	})
}
