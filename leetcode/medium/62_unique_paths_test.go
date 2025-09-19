package medium_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/leetcode/medium"
	"testing"
)

func TestUniquePaths(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, 3, medium.UniquePaths(3, 2))
	})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, 28, medium.UniquePaths(3, 7))
	})
}
