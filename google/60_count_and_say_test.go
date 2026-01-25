package google

import (
	"github.com/stretchr/testify/assert"
	"gocoder/leetcode/medium"
	"testing"
)

func TestCountAndSay(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, "1211", medium.CountAndSay(4))
	})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, "1", medium.CountAndSay(1))
	})

	//t.Run("t3", func(t *testing.T) {
	//	assert.Equal(t, "311311222113111231133211121312211231131112311211133112111312211213211312111322211231131122111213122112311311222112111331121113112221121113122113121113222112132113213221232112111312111213322112311311222113111221221113122112132113121113222112311311222", medium.CountAndSay(30))
	//})
}

func TestCountAndSayIter(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		assert.Equal(t, "1211", medium.CountAndSayIter(4))
	})

	t.Run("t2", func(t *testing.T) {
		assert.Equal(t, "1", medium.CountAndSayIter(1))
	})

	//t.Run("t3", func(t *testing.T) {
	//	assert.Equal(t, "311311222113111231133211121312211231131112311211133112111312211213211312111322211231131122111213122112311311222112111331121113112221121113122113121113222112132113213221232112111312111213322112311311222113111221221113122112132113121113222112311311222", medium.CountAndSay(30))
	//})
}
