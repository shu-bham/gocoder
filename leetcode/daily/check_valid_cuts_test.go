package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	leetcode "gocoder/leetcode/daily"
	"testing"
)

func TestCheckValidCuts(t *testing.T) {

	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, true, leetcode.CheckValidCuts(5, [][]int{{1, 0, 5, 2}, {0, 2, 2, 4}, {3, 2, 5, 3}, {0, 4, 4, 5}}))
	})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, true, leetcode.CheckValidCuts(4, [][]int{{0, 0, 1, 1}, {2, 0, 3, 4}, {0, 2, 2, 3}, {3, 0, 4, 3}}))
	})

	t.Run("t3", func(t *testing.T) {
		assert.Equal(t, false, leetcode.CheckValidCuts(4, [][]int{{0, 2, 2, 4}, {1, 0, 3, 2}, {2, 2, 3, 4}, {3, 0, 4, 2}, {3, 2, 4, 4}}))
	})
}
