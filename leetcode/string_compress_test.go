package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/leetcode"
	"testing"
)

func TestCompress(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
		len := leetcode.Compress(chars)
		assert.Equal(t, 6, len)
	})

	t.Run("t2", func(t *testing.T) {
		chars := []byte{'a'}
		len := leetcode.Compress(chars)
		assert.Equal(t, 1, len)
	})

	t.Run("t3", func(t *testing.T) {
		chars := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
		len := leetcode.Compress(chars)
		assert.Equal(t, 4, len)
	})
}

func TestCompressV2(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
		len := leetcode.CompressV2(chars)
		assert.Equal(t, 6, len)
	})

	t.Run("t2", func(t *testing.T) {
		chars := []byte{'a'}
		len := leetcode.CompressV2(chars)
		assert.Equal(t, 1, len)
	})

	t.Run("t3", func(t *testing.T) {
		chars := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
		len := leetcode.CompressV2(chars)
		assert.Equal(t, 4, len)
	})
}
