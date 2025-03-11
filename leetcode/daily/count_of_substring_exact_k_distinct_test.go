package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountOfSubstringsExactK(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, 10, CountOfSubstringsExactK("abcabc", 3))
	})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, 3, CountOfSubstringsExactK("aaacb", 3))
	})

	t.Run("t3", func(t *testing.T) {
		assert.Equal(t, 1, CountOfSubstringsExactK("abc", 3))
	})

	t.Run("single character string", func(t *testing.T) {
		assert.Equal(t, 0, CountOfSubstringsExactK("a", 2))
	})

	t.Run("all identical characters", func(t *testing.T) {
		assert.Equal(t, 0, CountOfSubstringsExactK("aaaa", 3))
	})

	t.Run("longer string", func(t *testing.T) {
		assert.Equal(t, 23, CountOfSubstringsExactK("abacababc", 3))
	})
}

func TestCountOfSubstringsExactKV2(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, 10, CountOfSubstringsExactKV2("abcabc", 3))
	})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, 3, CountOfSubstringsExactKV2("aaacb", 3))
	})

	t.Run("t3", func(t *testing.T) {
		assert.Equal(t, 1, CountOfSubstringsExactKV2("abc", 3))
	})

	t.Run("single character string", func(t *testing.T) {
		assert.Equal(t, 0, CountOfSubstringsExactKV2("a", 2))
	})

	t.Run("all identical characters", func(t *testing.T) {
		assert.Equal(t, 0, CountOfSubstringsExactKV2("aaaa", 3))
	})

	t.Run("longer string", func(t *testing.T) {
		assert.Equal(t, 23, CountOfSubstringsExactKV2("abacababc", 3))
	})
}
