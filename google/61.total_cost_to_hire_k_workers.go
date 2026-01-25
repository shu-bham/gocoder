package google

import (
	"container/heap"
	"math"
)

// brute
func totalCost(costs []int, k int, candidates int) int64 {
	n := len(costs)
	hired := make([]bool, n)
	var total int64 = 0

	// Assuming the problem constraint k <= len(costs) holds.
	if n == 0 {
		return 0
	}

	for i := 0; i < k; i++ {
		// Find the minimum from the left pool of 'candidates' unhired workers.
		minCost1, minIndex1 := math.MaxInt, math.MaxInt
		count1 := 0
		for j := 0; j < n && count1 < candidates; j++ {
			if !hired[j] {
				if costs[j] < minCost1 {
					minCost1 = costs[j]
					minIndex1 = j
				} else if costs[j] == minCost1 {
					if j < minIndex1 {
						minIndex1 = j
					}
				}
				count1++
			}
		}

		// Find the minimum from the right pool of 'candidates' unhired workers.
		minCost2, minIndex2 := math.MaxInt, math.MaxInt
		count2 := 0
		for j := n - 1; j >= 0 && count2 < candidates; j-- {
			if !hired[j] {
				if costs[j] < minCost2 {
					minCost2 = costs[j]
					minIndex2 = j
				} else if costs[j] == minCost2 {
					if j < minIndex2 {
						minIndex2 = j
					}
				}
				count2++
			}
		}

		if minCost1 <= minCost2 {
			total += int64(minCost1)
			hired[minIndex1] = true
		} else {
			total += int64(minCost2)
			hired[minIndex2] = true
		}
	}

	return total
}

type worker struct {
	cost, pos int
}
type workerHeap []worker

func (w workerHeap) Len() int {
	return len(w)
}

func (w workerHeap) Less(i, j int) bool {
	if w[i].cost == w[j].cost {
		return w[i].pos < w[j].pos
	} else {
		return w[i].cost < w[j].cost
	}
}

func (w workerHeap) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func (w *workerHeap) Push(x any) {
	*w = append(*w, x.(worker))
}

func (w *workerHeap) Pop() any {
	old := *w
	n := old.Len()
	pop := old[n-1]
	*w = old[:n-1]
	return pop
}

func (w workerHeap) Peek() any {
	if w.Len() == 0 {
		return nil
	}
	return w[0]
}

func totalCostOpt(costs []int, k int, candidates int) int64 {
	var total int64 = 0
	leftHeap, rightHeap := &workerHeap{}, &workerHeap{}
	heap.Init(leftHeap)
	heap.Init(rightHeap)

	leftIndex, rightIndex := 0, len(costs)-1

	for leftHeap.Len() < candidates && leftIndex <= rightIndex {
		heap.Push(leftHeap, worker{costs[leftIndex], leftIndex})
		leftIndex++
	}

	for rightHeap.Len() < candidates && leftIndex <= rightIndex {
		heap.Push(rightHeap, worker{costs[rightIndex], rightIndex})
		rightIndex--
	}

	for k > 0 {
		choseLeft := false
		if leftHeap.Len() > 0 && rightHeap.Len() > 0 {
			c1 := leftHeap.Peek().(worker)
			c2 := rightHeap.Peek().(worker)

			if c1.cost == c2.cost {
				if c1.pos < c2.pos {
					choseLeft = true
				}
			} else if c1.cost < c2.cost {
				choseLeft = true
			}
		} else if leftHeap.Len() > 0 {
			choseLeft = true
		} else {
			choseLeft = false
		}

		if choseLeft {
			pop := heap.Pop(leftHeap)
			total += int64(pop.(worker).cost)
			if leftIndex <= rightIndex {
				heap.Push(leftHeap, worker{costs[leftIndex], leftIndex})
				leftIndex++
			}
		} else {
			pop := heap.Pop(rightHeap)
			total += int64(pop.(worker).cost)
			if leftIndex <= rightIndex {
				heap.Push(rightHeap, worker{costs[rightIndex], rightIndex})
				rightIndex--
			}
		}
		k--
	}
	return total
}
