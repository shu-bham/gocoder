# Analysis of Jump Game II

This document breaks down various approaches to solving the "Jump Game II" problem, progressing from a brute-force method to the most optimal greedy solution.

### Problem Statement

Given an array of non-negative integers `nums`, you are initially positioned at the first index. Each element in the array represents your maximum jump length at that position. Your goal is to reach the last index in the minimum number of jumps. It is guaranteed that the last index can be reached.

---

### Approach 1: Brute-Force Recursion

The most intuitive approach is to try every possible jump from every position recursively. We can define a function that calculates the minimum jumps from a given index `i` to the end.

**Logic:**
For a given index `i`, we can jump to any index `j` where `i < j <= i + nums[i]`. The minimum jumps from `i` would be `1 + min(jumps_from_j)`.

```go
package leetcode

import "math"

// jumpRecursive is a helper function for the brute-force approach.
func jumpRecursive(nums []int, index int) int {
	// Base case: If we are at or beyond the last index, no more jumps are needed.
	if index >= len(nums)-1 {
		return 0
	}

	minJumps := math.MaxInt32 -1 // Use a large value to represent infinity

	// Explore all possible jumps from the current index.
	for i := 1; i <= nums[index]; i++ {
		// Find the minimum jumps required from the next position.
		jumps := jumpRecursive(nums, index+i)
		if jumps+1 < minJumps {
			minJumps = jumps + 1
		}
	}

	return minJumps
}

// JumpBruteForce starts the recursive process from index 0.
func JumpBruteForce(nums []int) int {
	return jumpRecursive(nums, 0)
}
```

*   **Time Complexity: O(2^n)**. This is exponential because we re-compute the same subproblems repeatedly.
*   **Space Complexity: O(n)** for the recursion stack depth.
*   **Result:** This solution is too slow for most platforms and will result in a "Time Limit Exceeded" (TLE) error.

---

### Approach 2: Dynamic Programming (Top-Down with Memoization)

We can optimize the brute-force solution by storing the results of subproblems that we have already solved. We can use a `memo` array where `memo[i]` stores the minimum jumps needed to reach the end from index `i`.

**Logic:**
Before computing the jumps from an index `i`, we first check if `memo[i]` has already been computed. If it has, we return the stored value. Otherwise, we compute it, store it, and then return it.

```go
package leetcode

import "math"

// jumpMemo is the helper for the DP with memoization approach.
func jumpMemo(nums []int, index int, memo []int) int {
	// Base case: If we are at or beyond the last index, no more jumps are needed.
	if index >= len(nums)-1 {
		return 0
	}

	// If we have already computed the result for this index, return it.
	if memo[index] != 0 {
		return memo[index]
	}

	minJumps := math.MaxInt32 - 1

	// Explore all possible jumps from the current index.
	for i := 1; i <= nums[index]; i++ {
		jumps := jumpMemo(nums, index+i, memo)
		if jumps+1 < minJumps {
			minJumps = jumps + 1
		}
	}

	// Store the result before returning.
	memo[index] = minJumps
	return minJumps
}

// JumpDP starts the DP process.
func JumpDP(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	memo := make([]int, len(nums))
	return jumpMemo(nums, 0, memo)
}
```

*   **Time Complexity: O(n^2)**. Each subproblem `memo[i]` is solved only once. To solve for `i`, we might iterate up to `n` steps.
*   **Space Complexity: O(n)** for the recursion stack and the memoization array.
*   **Result:** This passes on most platforms but is not the most optimal solution.

---

### Approach 3: Greedy Algorithm (Optimal)

The problem can be solved in a single pass using a greedy strategy. We don't need to explore every possible jump; we only need to know the farthest point we can reach.

**Logic:**
We maintain `currentReach`, the maximum index we can reach with the current number of jumps, and `farthestReach`, the maximum index we can reach from any point within the `currentReach`.

1.  Iterate through the array from `i = 0` to `n-2`.
2.  At each position `i`, update `farthestReach` to be `max(farthestReach, i + nums[i])`. This tells us the best we can do from our current position.
3.  When our iterator `i` reaches the `currentReach`, it means we have exhausted all options for the current jump and we must take a new one. We increment our `jumps` counter and update `currentReach` to the `farthestReach` we have found so far.

This greedy choice works because we are always making a jump that lands us in a position with the potential to jump the farthest.

```go
package leetcode

// JumpGameII (renamed from JumpGreedy for clarity) solves the problem using the optimal greedy algorithm.
func JumpGameII(nums []int) int {
	// If there's only one element or it's empty, we're already at the end.
	if len(nums) <= 1 {
		return 0
	}

	jumps := 0
	currentReach := 0
	farthestReach := 0

	// We iterate from 0 to n-2. We don't need to consider a jump from the last element.
	for i := 0; i < len(nums)-1; i++ {
		// Update the farthest reach we can achieve.
		if i+nums[i] > farthestReach {
			farthestReach = i + nums[i]
		}

		// If we've reached the end of the current jump's reach, we must take another jump.
		if i == currentReach {
			jumps++
			currentReach = farthestReach

			// Optimization: if the new reach is already past the end, we can stop.
			if currentReach >= len(nums)-1 {
				break
			}
		}
	}

	return jumps
}
```

*   **Time Complexity: O(n)**. We iterate through the array once.
*   **Space Complexity: O(1)**. We only use a few variables to track our state.
*   **Result:** This is the most efficient and optimal solution.
