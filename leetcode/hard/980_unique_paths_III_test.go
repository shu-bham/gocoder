package hard_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/leetcode/hard"
	"testing"
)

func TestUniquePathsIII(t *testing.T) {

	t.Run("tc1", func(t *testing.T) {
		input := [][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 2, -1}}
		output := hard.UniquePathsIII(input)
		assert.Equal(t, 2, output)
	})

	t.Run("tc2", func(t *testing.T) {
		input := [][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 2}}
		output := hard.UniquePathsIII(input)
		assert.Equal(t, 4, output)
	})

	t.Run("tc3", func(t *testing.T) {
		input := [][]int{{0, 1}, {2, 0}}
		output := hard.UniquePathsIII(input)
		assert.Equal(t, 0, output)
	})
}
