package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	leetcode "gocoder/leetcode/daily"
	"testing"
)

func TestLongestNiceSubarray(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, 3, leetcode.LongestNiceSubarray([]int{1, 3, 8, 48, 10}))
	})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, 3, leetcode.LongestNiceSubarray([]int{1, 3, 8, 48, 10, 3, 8}))
	})

	t.Run("t3", func(t *testing.T) {
		assert.Equal(t, 2, leetcode.LongestNiceSubarray([]int{1, 3, 10, 3, 8}))
	})
}

func TestLongestNiceSubarrayV2(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, 3, leetcode.LongestNiceSubarrayV2([]int{1, 3, 8, 48, 10}))
	})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, 3, leetcode.LongestNiceSubarrayV2([]int{1, 3, 8, 48, 10, 3, 8}))
	})

	t.Run("t3", func(t *testing.T) {
		assert.Equal(t, 2, leetcode.LongestNiceSubarrayV2([]int{1, 3, 10, 3, 8}))
	})
}
