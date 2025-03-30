package leetcode

import (
	"container/heap"
	"math"
)

// Hard: https://leetcode.com/problems/apply-operations-to-maximize-score

func primeScore(n int) int {
	score := 0
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			score++
			for n%i == 0 {
				n = n / i
			}
		}
	}

	if n >= 2 {
		score++
	}

	return score
}

type Pair struct {
	value int
	index int
}

type MaxHeap []Pair

func (h MaxHeap) Len() int { return len(h) }

func (h MaxHeap) Less(i, j int) bool {
	if h[i].value == h[j].value {
		return h[i].index < h[j].index
	}
	return h[i].value > h[j].value
}

func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(Pair)) }

func (h *MaxHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

const MOD = 1_000_000_007

func MaximumScoreAfterApplyingOps(nums []int, k int) int {
	n := len(nums)
	primeScores := make([]int, n)
	for i, num := range nums {
		primeScores[i] = primeScore(num)
	}

	nextDominant := make([]int, n)
	prevDominant := make([]int, n)

	for i := range nextDominant {
		nextDominant[i] = n
	}

	for i := range prevDominant {
		prevDominant[i] = -1
	}

	// monotonic stack
	stack := []int{}
	for i := range nums {
		for len(stack) > 0 && primeScores[stack[len(stack)-1]] < primeScores[i] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			nextDominant[top] = i
		}
		if len(stack) > 0 {
			prevDominant[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	numOfSubarray := make([]int, n)
	for i := range nums {
		numOfSubarray[i] = (nextDominant[i] - i) * (i - prevDominant[i])
	}

	h := &MaxHeap{}
	heap.Init(h)
	for i, num := range nums {
		heap.Push(h, Pair{
			value: num,
			index: i,
		})
	}

	score := 1
	for k > 0 {
		pop := heap.Pop(h).(Pair)
		num := pop.value
		index := pop.index

		ops := min(k, numOfSubarray[index])
		score = (score * power(num, ops)) % MOD
		k -= ops
	}

	return score
}

func power(base, exponent int) int {
	res := 1
	b := base
	n := exponent
	for n > 0 {
		if n%2 == 1 {
			res = (res * b) % MOD
		}
		b = (b * b) % MOD
		n /= 2
	}
	return res
}
