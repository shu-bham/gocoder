package easy_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/leetcode/easy"
	"testing"
)

func TestNumOfLaserBeams(t *testing.T) {

	t.Run("tc1", func(t *testing.T) {
		assert.Equal(t, 8, easy.NumberOfBeams([]string{"011001", "000000", "010100", "001000"}))
	})

	t.Run("tc2", func(t *testing.T) {
		assert.Equal(t, 0, easy.NumberOfBeams([]string{"000", "111", "000"}))
	})
}

func TestNumOfLaserBeamsV2(t *testing.T) {

	t.Run("tc1", func(t *testing.T) {
		assert.Equal(t, 8, easy.NumberOfBeamsV2([]string{"011001", "000000", "010100", "001000"}))
	})

	t.Run("tc2", func(t *testing.T) {
		assert.Equal(t, 0, easy.NumberOfBeamsV2([]string{"000", "111", "000"}))
	})
}
