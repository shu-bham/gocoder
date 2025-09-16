package medium_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/leetcode/medium"
	"testing"
)

func TestSetZeroes(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		in := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
		medium.SetZeroes(in)
		assert.Equal(t, [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}, in)
	})

	t.Run("t2", func(t *testing.T) {
		in := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
		medium.SetZeroes(in)
		assert.Equal(t, [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}, in)
	})
}
