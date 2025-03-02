package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/leetcode"
	"testing"
)

func TestHIndex(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		citations := []int{3, 0, 6, 1, 5}
		hIndex := leetcode.HIndex(citations)
		assert.Equal(t, 3, hIndex)
	})
}

func TestHIndexV2(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		citations := []int{3, 0, 6, 1, 5}
		hIndex := leetcode.HIndexV2(citations)
		assert.Equal(t, 3, hIndex)
	})
}
