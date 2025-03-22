package misc_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/misc"
	"testing"
)

func TestTopologicalSortInDegree(t *testing.T) {

	t.Run("t1", func(t *testing.T) {
		topoSort, ok := misc.TopologicalSortInDegree(6, [][]int{
			{0, 1}, {1, 2}, {2, 3},
			{4, 5}, {5, 1}, {5, 2},
		})
		assert.True(t, ok)
		assert.Equal(t, []int{0, 4, 5, 1, 2, 3}, topoSort)
	})

	t.Run("t2", func(t *testing.T) {
		topoSort, ok := misc.TopologicalSortInDegree(3, [][]int{
			{0, 1}, {1, 2}, {2, 0},
		})
		assert.False(t, ok)
		assert.Equal(t, []int(nil), topoSort)
	})
}

func TestTopologicalSortDFS(t *testing.T) {

	t.Run("t1", func(t *testing.T) {
		topoSort, ok := misc.TopologicalSortDFS(6, [][]int{
			{0, 1}, {1, 2}, {2, 3},
			{4, 5}, {5, 1}, {5, 2},
		})
		assert.True(t, ok)
		assert.Equal(t, []int{4, 5, 0, 1, 2, 3}, topoSort)
	})

	t.Run("t2", func(t *testing.T) {
		topoSort, ok := misc.TopologicalSortDFS(3, [][]int{
			{0, 1}, {1, 2}, {2, 0},
		})
		assert.False(t, ok)
		assert.Equal(t, []int(nil), topoSort)
	})
}
