package misc_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/misc"
	"testing"
)

func TestBFS(t *testing.T) {
	t.Run("t1", func(t *testing.T) {
		bfs := misc.BFS([][]int{
			{0, 1},
			{0, 2},
			{2, 3},
			{1, 3},
			{1, 4},
			{4, 5},
			{4, 6}}, 0)
		assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6}, bfs)
	})

	t.Run("t2", func(t *testing.T) {
		bfs := misc.BFS([][]int{
			{1, 2},
			{1, 5},
			{2, 5},
			{2, 3},
			{3, 5},
			{4, 5},
			{3, 4}}, 1)
		assert.Equal(t, []int{1, 2, 5, 3, 4}, bfs)
	})

	t.Run("single node", func(t *testing.T) {
		bfs := misc.BFS([][]int{}, 0)
		assert.Equal(t, []int{0}, bfs)
	})

	t.Run("disconnected graph", func(t *testing.T) {
		bfs := misc.BFS([][]int{
			{0, 1},
			{1, 2},
			{3, 4},
			{4, 5}}, 0)
		assert.Equal(t, []int{0, 1, 2}, bfs)
	})

	t.Run("graph with cycle", func(t *testing.T) {
		bfs := misc.BFS([][]int{
			{0, 1},
			{1, 2},
			{2, 0},
			{2, 3},
			{3, 4}}, 0)
		assert.Equal(t, []int{0, 1, 2, 3, 4}, bfs)
	})

	t.Run("start from different node", func(t *testing.T) {
		bfs := misc.BFS([][]int{
			{0, 1},
			{0, 2},
			{1, 3},
			{2, 4},
			{4, 5}}, 2)
		assert.Equal(t, []int{2, 0, 4, 1, 5, 3}, bfs)
	})
}
