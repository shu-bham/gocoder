package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/leetcode"
	"testing"
)

func TestCountPrimes(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, 4, leetcode.CountPrimes(10))
	})
}

func TestCountPrimesV2(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, 4, leetcode.CountPrimesV2(10))
	})
}
