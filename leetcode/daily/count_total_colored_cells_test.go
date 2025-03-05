package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	leetcode "gocoder/leetcode/daily"
	"testing"
)

func TestColoredCells(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, int64(1), leetcode.ColoredCells(1))
	})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, int64(5), leetcode.ColoredCells(2))
	})

	t.Run("t3", func(t *testing.T) {
		assert.Equal(t, int64(13), leetcode.ColoredCells(3))
	})

	t.Run("t4", func(t *testing.T) {
		assert.Equal(t, int64(25), leetcode.ColoredCells(4))
	})

	t.Run("t5", func(t *testing.T) {
		assert.Equal(t, int64(41), leetcode.ColoredCells(5))
	})
}

func TestColoredCellsV2(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, int64(1), leetcode.ColoredCellsV2(1))
	})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, int64(5), leetcode.ColoredCellsV2(2))
	})

	t.Run("t3", func(t *testing.T) {
		assert.Equal(t, int64(13), leetcode.ColoredCellsV2(3))
	})

	t.Run("t4", func(t *testing.T) {
		assert.Equal(t, int64(25), leetcode.ColoredCellsV2(4))
	})

	t.Run("t5", func(t *testing.T) {
		assert.Equal(t, int64(41), leetcode.ColoredCellsV2(5))
	})
}
