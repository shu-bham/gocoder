package google

import "container/heap"

type pq []int

func (p pq) Len() int {
	return len(p)
}

func (p pq) Less(i, j int) bool {
	// We want a max-heap, so we use >
	return p[i] > p[j]
}

func (p pq) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *pq) Push(x any) {
	*p = append(*p, x.(int))
}

func (p *pq) Pop() any {
	old := *p
	n := old.Len()
	item := old[n-1]
	*p = old[:n-1]
	return item
}

func FindKthLargest(nums []int, k int) int {
	// Create a slice of type pq (which is []int) and copy the data.
	// This prevents modifying the original input slice.
	maxHeapData := make(pq, len(nums))
	copy(maxHeapData, nums)

	// heap.Init on a pre-populated slice is an O(N) operation.
	// We pass a pointer because the heap functions modify the slice's length.
	heap.Init(&maxHeapData)

	// Pop k-1 elements to get to the kth largest.
	for i := 0; i < k-1; i++ {
		heap.Pop(&maxHeapData)
	}

	// The top of the heap is now the kth largest element.
	return heap.Pop(&maxHeapData).(int)
}
