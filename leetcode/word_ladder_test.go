package leetcode_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/leetcode"
	"testing"
)

func TestLadderLength(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		begin, end, list := "hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}
		ladderLength := leetcode.LadderLength(begin, end, list)
		assert.Equal(t, 5, ladderLength)
	})

	t.Run("t2", func(t *testing.T) {
		begin, end, list := "hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}
		ladderLength := leetcode.LadderLength(begin, end, list)
		assert.Equal(t, 0, ladderLength)
	})

	t.Run("t3", func(t *testing.T) {
		begin, end, list := "kiss", "tusk", []string{"miss", "dusk", "kiss", "musk", "tusk", "diss", "disk", "sang", "ties", "muss"}
		ladderLength := leetcode.LadderLength(begin, end, list)
		assert.Equal(t, 5, ladderLength)
	})

	t.Run("t4", func(t *testing.T) {
		begin, end, list := "qa", "sq", []string{"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb", "sb", "kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br", "ti", "ba", "to", "ra", "fa", "yo", "ow", "sn", "ya", "cr", "po", "fe", "ho", "ma", "re", "or", "rn", "au", "ur", "rh", "sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if", "pb", "ge", "th", "pm", "rb", "sh", "co", "ga", "li", "ha", "hz", "no", "bi", "di", "hi", "qa", "pi", "os", "uh", "wm", "an", "me", "mo", "na", "la", "st", "er", "sc", "ne", "mn", "mi", "am", "ex", "pt", "io", "be", "fm", "ta", "tb", "ni", "mr", "pa", "he", "lr", "sq", "ye"}
		ladderLength := leetcode.LadderLength(begin, end, list)
		assert.Equal(t, 5, ladderLength)
	})
}
