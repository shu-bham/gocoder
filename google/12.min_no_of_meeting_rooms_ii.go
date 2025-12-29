package google

import (
	"container/heap"
	"sort"
)

type Interval struct {
	Start int
	End   int
}

type minHeap []Interval

func (p minHeap) Len() int {
	return len(p)
}

func (p minHeap) Less(i, j int) bool {
	return p[i].End < p[j].End
}

func (p minHeap) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *minHeap) Push(x any) {
	*p = append(*p, x.(Interval))
}

func (p *minHeap) Pop() any {
	old := *p
	n := old.Len()
	item := old[n-1]
	*p = old[:n-1]
	return item
}

func (p minHeap) Peek() any {
	return p[0]
}

// using minHeap
func MinMeetingRooms(intervals []Interval) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	pq := &minHeap{}
	heap.Init(pq)

	ans := 0
	for _, interval := range intervals {

		for pq.Len() > 0 && interval.Start >= pq.Peek().(Interval).End {
			heap.Pop(pq)
		}

		heap.Push(pq, Interval{
			Start: interval.Start,
			End:   interval.End,
		})
		ans = max(ans, pq.Len())
	}

	return ans
}

// using linesweep algo
func MinMeetingRoomsv2(intervals []Interval) int {
	maxEnd := 0
	for _, interval := range intervals {
		maxEnd = max(interval.End, maxEnd)
	}

	arr := make([]int, maxEnd+1)

	for _, interval := range intervals {
		arr[interval.Start]++
		arr[interval.End]--
	}

	ans := 0
	prev := 0
	for _, val := range arr {
		prev += val
		ans = max(ans, prev)
	}
	return ans
}
