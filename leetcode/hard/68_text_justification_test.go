package hard_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/leetcode/hard"
	"testing"
)

func TestFullJustify(t *testing.T) {

	t.Run("t1", func(t *testing.T) {
		maxWidth := 16
		words := []string{"This", "is", "an", "example", "of", "text", "justification."}
		strings := hard.FullJustify(words, maxWidth)
		exp := []string{"This    is    an", "example  of text", "justification.  "}
		assert.Equal(t, exp, strings)
	})

	t.Run("t2", func(t *testing.T) {
		maxWidth := 16
		words := []string{"What", "must", "be", "acknowledgment", "shall", "be"}
		strings := hard.FullJustify(words, maxWidth)
		exp := []string{"What   must   be", "acknowledgment  ", "shall be        "}
		assert.Equal(t, exp, strings)
	})

	t.Run("t3", func(t *testing.T) {
		maxWidth := 20
		words := []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
		strings := hard.FullJustify(words, maxWidth)
		exp := []string{"Science  is  what we", "understand      well", "enough to explain to", "a  computer.  Art is", "everything  else  we", "do                  "}
		assert.Equal(t, exp, strings)
	})
}
