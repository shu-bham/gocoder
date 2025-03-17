package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/leetcode"
	"testing"
)

func TestFindJudge(t *testing.T) {

	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, 2, leetcode.FindJudge(2, [][]int{{1, 2}}))
	})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, 3, leetcode.FindJudge(3, [][]int{{1, 3}, {2, 3}}))
	})

	t.Run("t3", func(t *testing.T) {
		assert.Equal(t, -1, leetcode.FindJudge(3, [][]int{{1, 3}, {2, 3}, {3, 1}}))
	})

	t.Run("t4", func(t *testing.T) {
		assert.Equal(t, -1, leetcode.FindJudge(3, [][]int{{1, 2}, {2, 3}}))
	})
}
