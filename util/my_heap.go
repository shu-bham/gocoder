package util

type Heap struct {
	Arr []int
}

type IMinHeap interface {
	GetMin() int
	Insert(v int)
	Delete(v int)
	ExtractMin() int
	DecreaseKey(index int, newVal int)
}

func BuildMinHeap(arr []int) Heap {
	h := Heap{arr}
	n := len(h.Arr)
	for i := (n / 2) - 1; i >= 0; i-- {
		h.minHeapify(i)
	}
	return h
}

func (h *Heap) minHeapify(index int) {
	leftChildInd := 2*index + 1
	rightChildInd := 2*index + 2
	n := len(h.Arr)

	smallest := index
	if leftChildInd < n && h.Arr[leftChildInd] < h.Arr[smallest] {
		smallest = leftChildInd
	}
	if rightChildInd < n && h.Arr[rightChildInd] < h.Arr[smallest] {
		smallest = rightChildInd

	}
	if smallest != index {
		h.swapEl(index, smallest)
		h.minHeapify(smallest)
	}
}

func (h *Heap) swapEl(i int, j int) {
	h.Arr[i], h.Arr[j] = h.Arr[j], h.Arr[i]
}

func (h *Heap) GetMin() int {
	if len(h.Arr) == 0 {
		return -1
	}
	return h.Arr[0]
}

func (h *Heap) Insert(v int) {
	h.Arr = append(h.Arr, v)
	i := len(h.Arr) - 1

	parentInd := func(i int) int { return (i - 1) / 2 }

	for i > 0 && h.Arr[parentInd(i)] > h.Arr[i] {
		h.swapEl(i, parentInd(i))
		i = parentInd(i)
	}

}

func (h *Heap) Delete(v int) {
	n := len(h.Arr)
	if n == 0 {
		return
	}

	// Find index of element to delete
	index := -1
	for i, val := range h.Arr {
		if val == v {
			index = i
			break
		}
	}
	if index == -1 {
		return // Value not found
	}

	h.swapEl(index, n-1)
	h.Arr = h.Arr[:n-1]
	if index < len(h.Arr) {
		h.minHeapify(index)
	}
}

func (h *Heap) DecreaseKey(index int, newVal int) {
	if index < 0 || index >= len(h.Arr) {
		return
	}
	h.Arr[index] = newVal
	parentInd := func(i int) int { return (i - 1) / 2 }
	for index > 0 && h.Arr[parentInd(index)] > h.Arr[index] {
		h.swapEl(index, parentInd(index))
		index = parentInd(index)
	}
}

func (h *Heap) ExtractMin() int {
	n := len(h.Arr)
	if n == 0 {
		return -1
	}
	if n == 1 {
		root := h.Arr[0]
		h.Arr = nil
		return root
	}
	root := h.Arr[0]
	h.Arr[0] = h.Arr[n-1]
	h.Arr = h.Arr[:n-1]
	h.minHeapify(0)
	return root

}
