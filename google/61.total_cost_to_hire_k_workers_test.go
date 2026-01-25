package google

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// This file contains test cases for the totalCost and totalCostOpt functions.
// The structure is designed to be illustrative of how one might approach testing
// in a technical interview setting.

/*
 ### Technical Interview: Problem Solving and Testing Walkthrough

 #### 1. Clarifying Questions (Understanding the Problem)

 Before jumping into code, it's crucial to clarify the requirements and constraints.
 This demonstrates a methodical approach and avoids misunderstandings.

 - **Input Constraints:**
   - What is the range for `len(costs)`? (e.g., 1 to 10^5)
   - What is the range for each `costs[i]`? (e.g., 1 to 10^5)
   - What are the ranges for `k` and `candidates`? (e.g., `1 <= k, candidates <= len(costs)`)

 - **Behavioral Questions:**
   - What happens if `costs` is empty? (Should return 0)
   - How should ties in cost be handled? (The problem specifies choosing the worker with the lower index).
   - Can `k` be greater than the number of workers? (Assume `k <= len(costs)` based on typical problem constraints).

 #### 2. Algorithm Design (Devising a Solution)

 - **Problem:** We need to hire `k` workers with the minimum possible cost. The twist is that in each hiring session, we can only choose from two pools: the first `candidates` workers and the last `candidates` workers from the current list of available workers.

 - **Brute-Force Approach:**
   - For each of the `k` hires:
     1. Create two lists: one with the first `candidates` available workers and one with the last `candidates`.
     2. Find the minimum cost worker from these two lists, handling ties by index.
     3. Add the cost to the total and mark the worker as hired.
   - **Complexity:** O(k * candidates), which can be slow if `k` and `candidates` are large. The `totalCost` function in the source file implements this approach, serving as a good baseline and for validating the optimized solution.

 - **Optimized Approach (Using Heaps):**
   - The brute-force method repeatedly scans the candidate pools. We can do better. A min-heap is a perfect data structure for efficiently finding the minimum element in a collection.
   - **Algorithm:**
     1. Create two min-heaps: `leftHeap` for the candidates at the front and `rightHeap` for candidates at the back.
     2. Initialize the heaps: Populate `leftHeap` with the first `candidates` workers and `rightHeap` with the last `candidates` workers. Keep track of the boundaries with `leftIndex` and `rightIndex`.
     3. Loop `k` times:
        a. Look at the top (minimum) element of both heaps.
        b. Choose the worker with the lower cost. If costs are tied, choose the one with the smaller index (which `leftHeap` would represent).
        c. Add the chosen worker's cost to `total`.
        d. Remove the chosen worker from its heap.
        e. If there are more workers available between `leftIndex` and `rightIndex`, add the next available worker to the heap from which the worker was just taken.
   - **Complexity:** O(candidates) to build the initial heaps and O(k * log(candidates)) for the hiring process. This is much more efficient. The `totalCostOpt` function implements this.

*/

func TestTotalCostSolutions(t *testing.T) {
	testCases := []struct {
		name       string
		costs      []int
		k          int
		candidates int
		expected   int64
	}{
		{
			name:       "LeetCode Example 1",
			costs:      []int{17, 12, 10, 2, 7, 2, 11, 20, 8},
			k:          3,
			candidates: 4,
			expected:   11,
		},
		{
			name:       "LeetCode Example 2",
			costs:      []int{1, 2, 4, 1},
			k:          3,
			candidates: 3,
			expected:   4,
		},
		{
			name:       "All workers are candidates",
			costs:      []int{10, 1, 1, 10},
			k:          2,
			candidates: 4,
			expected:   2,
		},
		{
			name:       "Hiring all workers",
			costs:      []int{5, 4, 3, 2, 1},
			k:          5,
			candidates: 2,
			expected:   15,
		},
		{
			name:       "Single worker",
			costs:      []int{100},
			k:          1,
			candidates: 1,
			expected:   100,
		},
		{
			name:       "Empty costs array",
			costs:      []int{},
			k:          0,
			candidates: 0,
			expected:   0,
		},
		{
			name:       "Candidates greater than half",
			costs:      []int{6, 1, 2, 7, 3, 5, 4},
			k:          3,
			candidates: 5,
			expected:   6, // 1 + 2 + 3
		},
		{
			name:       "Tie in cost, choose smaller index",
			costs:      []int{10, 5, 20, 5, 15},
			k:          2,
			candidates: 2,
			expected:   10, // 5 (index 1) + 5 (index 3)
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name+"_BruteForce", func(t *testing.T) {
			result := totalCost(tc.costs, tc.k, tc.candidates)
			assert.Equal(t, tc.expected, result, "Brute-force solution failed")
		})
	}

	for _, tc := range testCases {
		t.Run(tc.name+"_Optimized", func(t *testing.T) {
			result := totalCostOpt(tc.costs, tc.k, tc.candidates)
			assert.Equal(t, tc.expected, result, "Optimized solution failed")
		})
	}
}
