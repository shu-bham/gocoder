package google

import "container/heap"

type pq []int

func (p pq) Len() int {
	return len(p)
}

func (p pq) Less(i, j int) bool {
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
	maxHeap := &pq{}
	heap.Init(maxHeap)

	for _, num := range nums {
		heap.Push(maxHeap, num)
	}

	var ans int
	for j := 1; j <= k; j++ {
		ans = heap.Pop(maxHeap).(int)
	}
	return ans

}
