package misc_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/misc"
	"math"
	"testing"
)

func TestShortestPathFromSrc(t *testing.T) {
	t.Run("empty graph", func(t *testing.T) {
		assert.Equal(t, []int{0, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt}, misc.ShortestPathFromSrc(5, 0, [][]int{}))
	})

	t.Run("single edge", func(t *testing.T) {
		assert.Equal(t, []int{0, 3, math.MaxInt, math.MaxInt, math.MaxInt}, misc.ShortestPathFromSrc(5, 0, [][]int{{0, 1, 3}}))
	})

	t.Run("multiple edges", func(t *testing.T) {
		assert.Equal(t, []int{0, 4, 12, 19, 21}, misc.ShortestPathFromSrc(5, 0, [][]int{{0, 1, 4}, {1, 2, 8}, {2, 3, 7}, {3, 4, 2}}))
	})

	t.Run("graph with cycle", func(t *testing.T) {
		assert.Equal(t, []int{0, 2, 5, 5, 3}, misc.ShortestPathFromSrc(5, 0, [][]int{{0, 1, 2}, {1, 2, 3}, {2, 3, 3}, {3, 4, 2}, {4, 1, 1}}))
	})
}
