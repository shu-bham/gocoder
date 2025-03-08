package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	leetcode "gocoder/leetcode/daily"
	"testing"
)

func TestClosestPrimes(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, []int{11, 13}, leetcode.ClosestPrimes(10, 19))
		assert.Equal(t, []int{-1, -1}, leetcode.ClosestPrimes(4, 6))
		assert.Equal(t, []int{1019, 1021}, leetcode.ClosestPrimes(1000, 10000))
	})
}
