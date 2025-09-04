package hard_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/leetcode/hard"
	"testing"
)

func TestMaxScoreWords(t *testing.T) {

	t.Run("tc1", func(t *testing.T) {
		words := []string{"dog", "cat", "dad", "good"}
		letters := []byte{'a', 'a', 'c', 'd', 'd', 'd', 'g', 'o', 'o'}
		score := []int{1, 0, 9, 5, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		res := hard.MaxScoreWords(words, letters, score)
		assert.Equal(t, 23, res)
	})

	t.Run("tc2", func(t *testing.T) {
		words := []string{"xxxz", "ax", "bx", "cx"}
		letters := []byte{'z', 'a', 'b', 'c', 'x', 'x', 'x'}
		score := []int{4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 10}
		res := hard.MaxScoreWords(words, letters, score)
		assert.Equal(t, 27, res)
	})

	t.Run("tc3", func(t *testing.T) {
		words := []string{"leetcode"}
		letters := []byte{'l', 'e', 't', 'c', 'o', 'd'}
		score := []int{0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0}
		res := hard.MaxScoreWords(words, letters, score)
		assert.Equal(t, 0, res)
	})
}
