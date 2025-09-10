package hard_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/leetcode/hard"
	"testing"
)

func TestColorTheGrid(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		m, n, expected := 1, 1, 3
		assert.Equal(t, expected, hard.ColorTheGrid(m, n))
	})

	t.Run("2", func(t *testing.T) {
		m, n, expected := 1, 2, 6
		assert.Equal(t, expected, hard.ColorTheGrid(m, n))
	})

	t.Run("3", func(t *testing.T) {
		m, n, expected := 5, 5, 580986
		assert.Equal(t, expected, hard.ColorTheGrid(m, n))
	})

	t.Run("4", func(t *testing.T) {
		m, n, expected := 2, 3, 54
		assert.Equal(t, expected, hard.ColorTheGrid(m, n))
	})

	t.Run("5", func(t *testing.T) {
		m, n, expected := 4, 2, 162
		assert.Equal(t, expected, hard.ColorTheGrid(m, n))
	})
}
