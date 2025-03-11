package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	leetcode "gocoder/leetcode/daily"
	"testing"
)

func TestCountOfSubstringsExactK(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, 10, leetcode.CountOfSubstringsExactK("abcabc", 3))
	})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, 3, leetcode.CountOfSubstringsExactK("aaacb", 3))
	})

	t.Run("t3", func(t *testing.T) {
		assert.Equal(t, 1, leetcode.CountOfSubstringsExactK("abc", 3))
	})

	t.Run("single character string", func(t *testing.T) {
		assert.Equal(t, 0, leetcode.CountOfSubstringsExactK("a", 2))
	})

	t.Run("all identical characters", func(t *testing.T) {
		assert.Equal(t, 0, leetcode.CountOfSubstringsExactK("aaaa", 3))
	})

	t.Run("longer string", func(t *testing.T) {
		assert.Equal(t, 23, leetcode.CountOfSubstringsExactK("abacababc", 3))
	})
}

func TestCountOfSubstringsExactKV2(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, 10, leetcode.CountOfSubstringsExactKV2("abcabc", 3))
	})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, 3, leetcode.CountOfSubstringsExactKV2("aaacb", 3))
	})

	t.Run("t3", func(t *testing.T) {
		assert.Equal(t, 1, leetcode.CountOfSubstringsExactKV2("abc", 3))
	})

	t.Run("single character string", func(t *testing.T) {
		assert.Equal(t, 0, leetcode.CountOfSubstringsExactKV2("a", 2))
	})

	t.Run("all identical characters", func(t *testing.T) {
		assert.Equal(t, 0, leetcode.CountOfSubstringsExactKV2("aaaa", 3))
	})

	t.Run("longer string", func(t *testing.T) {
		assert.Equal(t, 23, leetcode.CountOfSubstringsExactKV2("abacababc", 3))
	})
}
