package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	leetcode "gocoder/leetcode/daily"
	"testing"
)

func TestCountOfSubstrings(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, int64(1), leetcode.CountOfSubstrings("aeiouq", 1))
	})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, int64(0), leetcode.CountOfSubstrings("aeiou", 1))
	})
	t.Run("t3", func(t *testing.T) {
		assert.Equal(t, int64(0), leetcode.CountOfSubstrings("aeioqq", 1))
	})
	t.Run("t4", func(t *testing.T) {
		assert.Equal(t, int64(3), leetcode.CountOfSubstrings("ieaouqqieaouqq", 1))
	})
	t.Run("t5", func(t *testing.T) {
		assert.Equal(t, int64(3), leetcode.CountOfSubstrings("iqeaouqi", 2))
	})
}

func TestCountOfSubstringsV2(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, int64(1), leetcode.CountOfSubstringsV2("aeiouq", 1))
	})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, int64(0), leetcode.CountOfSubstringsV2("aeiou", 1))
	})
	t.Run("t3", func(t *testing.T) {
		assert.Equal(t, int64(0), leetcode.CountOfSubstringsV2("aeioqq", 1))
	})
	t.Run("t4", func(t *testing.T) {
		assert.Equal(t, int64(3), leetcode.CountOfSubstringsV2("ieaouqqieaouqq", 1))
	})
	t.Run("t5", func(t *testing.T) {
		assert.Equal(t, int64(3), leetcode.CountOfSubstringsV2("iqeaouqi", 2))
	})
}
