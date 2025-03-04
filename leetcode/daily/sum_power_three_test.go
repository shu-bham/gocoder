package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	leetcode "gocoder/leetcode/daily"
	"testing"
)

func TestCheckPowersOfThree(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, true, leetcode.CheckPowersOfThree(12))
	})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, true, leetcode.CheckPowersOfThree(91))
	})

	t.Run("t3", func(t *testing.T) {
		assert.Equal(t, false, leetcode.CheckPowersOfThree(21))
	})

	t.Run("t4", func(t *testing.T) {
		assert.Equal(t, true, leetcode.CheckPowersOfThree(1))
	})
}
