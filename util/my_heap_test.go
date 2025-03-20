package util_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/util"
	"testing"
)

func TestBuildMinHeap(t *testing.T) {
	t.Run("single element", func(t *testing.T) {
		assert.Equal(t, []int{1}, util.BuildMinHeap([]int{1}).Arr)
	})

	t.Run("already a min-heap", func(t *testing.T) {
		assert.Equal(t, []int{5, 10, 8}, util.BuildMinHeap([]int{10, 5, 8}).Arr)
	})

	t.Run("descending order input", func(t *testing.T) {
		assert.Equal(t, []int{5, 6, 13, 9, 15, 17}, util.BuildMinHeap([]int{17, 15, 13, 9, 6, 5}).Arr)
	})

	t.Run("empty input", func(t *testing.T) {
		assert.Equal(t, []int{}, util.BuildMinHeap([]int{}).Arr)
	})

	t.Run("duplicate values", func(t *testing.T) {
		assert.Equal(t, []int{3, 3, 5, 5, 9, 8, 8}, util.BuildMinHeap([]int{5, 3, 8, 3, 9, 5, 8}).Arr)
	})

	t.Run("already sorted input", func(t *testing.T) {
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, util.BuildMinHeap([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}).Arr)
	})
}

func TestHeapOperations(t *testing.T) {
	t.Run("GetMin on non-empty heap", func(t *testing.T) {
		h := util.BuildMinHeap([]int{10, 5, 8})
		assert.Equal(t, 5, h.GetMin())
	})

	t.Run("Insert maintains min-heap property", func(t *testing.T) {
		h := util.BuildMinHeap([]int{10, 5, 8})
		h.Insert(3)
		assert.Equal(t, 3, h.GetMin())
	})

	t.Run("Insert larger element keeps min unchanged", func(t *testing.T) {
		h := util.BuildMinHeap([]int{10, 5, 8})
		h.Insert(15)
		assert.Equal(t, 5, h.GetMin())
	})

	t.Run("ExtractMin removes and returns min element", func(t *testing.T) {
		h := util.BuildMinHeap([]int{10, 5, 8})
		min := h.ExtractMin()
		assert.Equal(t, 5, min)
		assert.NotContains(t, h.Arr, 5)
	})

	t.Run("ExtractMin on single element heap", func(t *testing.T) {
		h := util.BuildMinHeap([]int{42})
		min := h.ExtractMin()
		assert.Equal(t, 42, min)
		assert.Empty(t, h.Arr)
	})

	t.Run("Delete element from heap", func(t *testing.T) {
		h := util.BuildMinHeap([]int{3, 5, 8, 10, 15})
		h.Delete(5)
		assert.NotContains(t, h.Arr, 5)
	})

	t.Run("Delete non-existing element does nothing", func(t *testing.T) {
		h := util.BuildMinHeap([]int{3, 5, 8, 10, 15})
		h.Delete(20)
		assert.Equal(t, []int{3, 5, 8, 10, 15}, h.Arr)
	})

	t.Run("DecreaseKey updates and maintains min-heap", func(t *testing.T) {
		h := util.BuildMinHeap([]int{10, 20, 15, 30, 40})
		h.DecreaseKey(2, 5)
		assert.Equal(t, 5, h.GetMin())
	})

	t.Run("DecreaseKey out of bounds does nothing", func(t *testing.T) {
		h := util.BuildMinHeap([]int{10, 20, 15, 30, 40})
		assert.Equal(t, []int{10, 20, 15, 30, 40}, h.Arr)
		h.DecreaseKey(-1, 5)
		h.DecreaseKey(10, 5)
		assert.Equal(t, []int{10, 20, 15, 30, 40}, h.Arr)
	})
}
