package google

import "container/heap"

type element struct {
	num  int
	freq int
}
type priorityQueue []element

func (p priorityQueue) Len() int {
	return len(p)
}

func (p priorityQueue) Less(i, j int) bool {
	return p[i].freq > p[j].freq
}

func (p priorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *priorityQueue) Push(x any) {
	*p = append(*p, x.(element))
}

func (p *priorityQueue) Pop() any {
	old := *p
	n := old.Len()
	item := old[n-1]
	*p = old[:n-1]
	return item
}

func TopKFrequent(nums []int, k int) []int {
	var ans []int
	if len(nums) == 0 {
		return ans
	}
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	pq := &priorityQueue{}
	heap.Init(pq)
	for k, v := range freqMap {
		heap.Push(pq, element{
			num:  k,
			freq: v,
		})
	}

	for i := 0; i < k; i++ {
		ans = append(ans, heap.Pop(pq).(element).num)
	}

	return ans
}
