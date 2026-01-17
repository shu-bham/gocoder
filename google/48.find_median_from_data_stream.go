package google

import (
	"container/heap"
)

type myHeap struct {
	nums      []int
	isMinHeap bool
}

func (m *myHeap) Len() int {
	return len(m.nums)
}

func (m *myHeap) Less(i, j int) bool {
	if m.isMinHeap {
		return m.nums[i] < m.nums[j]
	} else {
		return m.nums[i] > m.nums[j]
	}
}

func (m *myHeap) Swap(i, j int) {
	m.nums[i], m.nums[j] = m.nums[j], m.nums[i]
}

func (m *myHeap) Push(x any) {
	m.nums = append(m.nums, x.(int))
}

func (m *myHeap) Pop() any {
	n := m.Len()
	item := m.nums[n-1]
	m.nums = m.nums[:n-1]
	return item
}

func (m *myHeap) Peek() int {
	if m.Len() == 0 {
		panic("peek on empty heap")
	}
	return m.nums[0]
}

type MedianFinder struct {
	leftMaxHeap  *myHeap
	rightMinHeap *myHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		leftMaxHeap: &myHeap{
			nums:      make([]int, 0),
			isMinHeap: false,
		},
		rightMinHeap: &myHeap{
			nums:      make([]int, 0),
			isMinHeap: true,
		},
	}
}

func (mf *MedianFinder) AddNum(num int) {
	if mf.rightMinHeap.Len() > 0 && num > mf.rightMinHeap.Peek() {
		heap.Push(mf.rightMinHeap, num)
	} else {
		heap.Push(mf.leftMaxHeap, num)
	}
	mf.BalanceHeaps()
	// fmt.Println("addNum::leftHeap:", num, mf.leftMaxHeap)
	// fmt.Println("addNum::rightHeap:", num, mf.rightMinHeap)
}

func (mf *MedianFinder) FindMedian() float64 {
	l1 := mf.leftMaxHeap.Len()
	l2 := mf.rightMinHeap.Len()
	if l1 == l2 {
		return float64(mf.leftMaxHeap.Peek()+mf.rightMinHeap.Peek()) / 2
	} else if l1 > l2 {
		return float64(mf.leftMaxHeap.Peek())
	} else {
		return float64(mf.rightMinHeap.Peek())
	}
}

func (mf *MedianFinder) BalanceHeaps() {
	leftHeapLen := mf.leftMaxHeap.Len()
	rightHeapLen := mf.rightMinHeap.Len()

	if abs(leftHeapLen-rightHeapLen) < 2 {
		return
	}

	if leftHeapLen > rightHeapLen {
		pop := heap.Pop(mf.leftMaxHeap)
		heap.Push(mf.rightMinHeap, pop)
	} else {
		pop := heap.Pop(mf.rightMinHeap)
		heap.Push(mf.leftMaxHeap, pop)
	}

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
