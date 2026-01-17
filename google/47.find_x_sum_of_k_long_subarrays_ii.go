package google

import (
	"container/heap"
)

type node struct {
	val, freq int
}

type myPQ struct {
	nums        []node
	positionMap map[int]int
	isMinHeap   bool
}

func (m *myPQ) Len() int {
	return len(m.nums)
}

func (m *myPQ) Less(i, j int) bool {
	if m.isMinHeap {
		if m.nums[i].freq == m.nums[j].freq {
			return m.nums[i].val < m.nums[j].val
		}
		return m.nums[i].freq < m.nums[j].freq
	}
	if m.nums[i].freq == m.nums[j].freq {
		return m.nums[i].val > m.nums[j].val
	}
	return m.nums[i].freq > m.nums[j].freq
}

func (m *myPQ) Swap(i, j int) {
	m.nums[i], m.nums[j] = m.nums[j], m.nums[i]
	m.positionMap[m.nums[i].val] = i
	m.positionMap[m.nums[j].val] = j
}

func (m *myPQ) Push(x any) {
	newNode := x.(node)
	m.positionMap[newNode.val] = len(m.nums)
	m.nums = append(m.nums, newNode)
}

func (m *myPQ) Pop() any {
	n := len(m.nums)
	item := m.nums[n-1]
	m.nums = m.nums[:n-1]
	delete(m.positionMap, item.val)
	return item
}

func (m *myPQ) Peek() node {
	return m.nums[0]
}

func findXSum(nums []int, k int, x int) []int64 {
	topXHeap := &myPQ{
		nums:        make([]node, 0),
		positionMap: make(map[int]int),
		isMinHeap:   true,
	}

	restHeap := &myPQ{
		nums:        make([]node, 0),
		positionMap: make(map[int]int),
		isMinHeap:   false,
	}

	freqMap := make(map[int]int)
	sum := int64(0)
	var ans []int64

	better := func(a, b node) bool {
		if a.freq != b.freq {
			return a.freq > b.freq
		}
		return a.val > b.val
	}

	rebalance := func() {
		for topXHeap.Len() > x {
			el := heap.Pop(topXHeap).(node)
			sum -= int64(el.val * el.freq)
			heap.Push(restHeap, el)
		}

		for topXHeap.Len() < x && restHeap.Len() > 0 {
			el := heap.Pop(restHeap).(node)
			sum += int64(el.val * el.freq)
			heap.Push(topXHeap, el)
		}

		// Swap elements between heaps if needed
		for restHeap.Len() > 0 && topXHeap.Len() > 0 {
			topMin := topXHeap.Peek()
			restMax := restHeap.Peek()

			if better(restMax, topMin) {
				heap.Pop(topXHeap)
				heap.Pop(restHeap)

				sum -= int64(topMin.val * topMin.freq)
				sum += int64(restMax.val * restMax.freq)

				heap.Push(topXHeap, restMax)
				heap.Push(restHeap, topMin)
			} else {
				break
			}
		}
	}

	addValue := func(val int) {
		oldFreq := freqMap[val]
		freqMap[val]++
		newFreq := freqMap[val]

		if oldFreq > 0 {
			if pos, ok := topXHeap.positionMap[val]; ok {
				heap.Remove(topXHeap, pos)
				sum -= int64(val * oldFreq)
			} else if pos, ok := restHeap.positionMap[val]; ok {
				heap.Remove(restHeap, pos)
			}
		}

		// Add new node to restHeap
		newNode := node{val, newFreq}
		heap.Push(restHeap, newNode)
		rebalance()
	}

	removeValue := func(val int) {
		oldFreq := freqMap[val]
		freqMap[val]--
		newFreq := freqMap[val]

		// Remove old node
		if pos, ok := topXHeap.positionMap[val]; ok {
			heap.Remove(topXHeap, pos)
			sum -= int64(val * oldFreq)
		} else if pos, ok := restHeap.positionMap[val]; ok {
			heap.Remove(restHeap, pos)
		}

		if newFreq > 0 {
			newNode := node{val, newFreq}
			heap.Push(restHeap, newNode)
		} else {
			delete(freqMap, val)
		}

		rebalance()
	}

	for i, num := range nums {
		addValue(num)
		if i >= k {
			removeValue(nums[i-k])
		}
		if i >= k-1 {
			ans = append(ans, sum)
		}
	}
	return ans
}
