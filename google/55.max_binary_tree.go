/*
# Maximum Binary Tree

## Problem Description

Given an integer array `nums` with no duplicates, a maximum binary tree can be built recursively from `nums` using the following algorithm:
1. Create a root node whose value is the maximum value in `nums`.
2. Recursively build the left subtree on the subarray prefix to the left of the maximum value.
3. Recursively build the right subtree on the subarray suffix to the right of the maximum value.
Return the maximum binary tree built from `nums`.

## Tech Interview Approach Protocol

When faced with this problem in a technical interview, here's a structured approach to follow:

### 1. Clarification and Edge Cases
Before writing any code, it's crucial to clarify the requirements with the interviewer.

- **Confirm Input/Output:**
    - **Input:** "The input is an array of integers, correct?"
    - **Output:** "And the function should return a pointer to the root node of the constructed tree?"

- **Ask Clarifying Questions:**
    - **Uniqueness:** "The problem description mentions 'no duplicates'. I'll assume all elements are unique. Is that correct?"
    - **Value Constraints:** "Are there any constraints on the integer values? Can they be positive, negative, or zero?"
    - **Malformed Inputs:** "How should I handle a `nil` slice as input vs. an empty slice `[]`?"

- **Discuss Edge Cases:**
    - **Empty Array:** "If the input array is empty, should the function return `nil`?" (This is the logical outcome).
    - **Single Element Array:** "For an array with just one element, the function should return a single TreeNode with no children, right?"

### 2. Naive/Brute-Force Approach (Recursive Solution)

This is the most intuitive approach and directly follows the problem description.

- **Algorithm:**
    1. Find the maximum element in the current subarray.
    2. Create a new `TreeNode` with this maximum value. This will be the root of the current (sub)tree.
    3. Recursively call the function for the part of the array to the left of the maximum element and assign the result to the `Left` child of the root.
    4. Recursively call the function for the part of the array to the right of the maximum element and assign the result to the `Right` child of the root.

- **Complexity Analysis:**
    - **Time Complexity:** In the worst case (e.g., a sorted array), we scan the array to find the max `N` times for an array of size `N`, then `N-1`, `N-2`, etc. This leads to an O(N^2) time complexity.
    - **Space Complexity:** The recursion depth can go up to `N` in the worst case (a skewed tree from a sorted array), leading to O(N) space complexity for the recursion stack.

The `constructMaximumBinaryTree` function in this file implements this O(N^2) approach.

### 3. Optimized Approach (Using a Monotonic Stack)

After discussing the naive approach, the interviewer will likely ask for a more optimal solution. The O(N^2) complexity comes from repeatedly scanning for the maximum element. We can optimize this by processing the array in a single pass. A monotonic stack is a great data structure for this.

- **Algorithm:**
    1. Iterate through the `nums` array, and for each number, create a new `TreeNode`.
    2. Maintain a monotonic stack of `TreeNode` pointers. The stack will store nodes in decreasing order of their values.
    3. For each new `node`:
        a. While the stack is not empty and the value of the node at the top of the stack is less than the current number's value, pop from the stack. The last popped node (`prev`) was smaller, so the current `node` must be its parent. We set `node.Left = prev`.
        b. If the stack is not empty after the while loop, the node at the top of the stack has a value greater than the current `node`. This means the current `node` is the right child of that top-of-stack node. So, set `stack[top].Right = node`.
        c. Push the current `node` onto the stack.

- **Why it works:**
    - The `Left` child of any node `X` is the root of the maximum binary tree constructed from the subarray between the parent of `X` and `X`.
    - The `Right` child of any node `Y` is the root of the maximum binary tree constructed from the subarray that appears after `Y` and before the next element larger than `Y`.
    - The stack helps find the parent (the first element to the left that is larger) and the left child (the largest element to the right that is smaller) efficiently.

- **Complexity Analysis:**
    - **Time Complexity:** Each element is pushed and popped from the stack exactly once. This gives a linear time complexity of O(N).
    - **Space Complexity:** In the worst case (e.g., a sorted descending array), all nodes will be in the stack. This gives a space complexity of O(N).

The `constructMaximumBinaryTreeOptimised` function in this file implements this O(N) approach.

### 4. Coding and Testing
- Write clean, commented code for the optimized O(N) solution.
- Manually trace the execution with a simple example like `[3, 2, 1, 6, 0, 5]` to demonstrate your understanding and verify the logic.
- Discuss how the code handles the edge cases you identified earlier (empty and single-element arrays).

*/
package google

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode1 struct {
	Val   int
	Left  *TreeNode1
	Right *TreeNode1
}

func constructMaximumBinaryTree(nums []int) *TreeNode1 {
	return constructMaximumBinaryTreeHelper(nums, 0, len(nums)-1)
}

func constructMaximumBinaryTreeHelper(nums []int, start int, end int) *TreeNode1 {
	if start > end {
		return nil
	}
	maxElement, maxPos := math.MinInt, -1
	for i := start; i <= end; i++ {
		if nums[i] > maxElement {
			maxElement = nums[i]
			maxPos = i
		}
	}
	node := &TreeNode1{Val: maxElement}
	node.Left = constructMaximumBinaryTreeHelper(nums, start, maxPos-1)
	node.Right = constructMaximumBinaryTreeHelper(nums, maxPos+1, end)
	return node
}

func constructMaximumBinaryTreeOptimised(nums []int) *TreeNode1 {
	stack := make([]*TreeNode1, 0)

	for _, num := range nums {
		node := &TreeNode1{Val: num}
		var prev *TreeNode1

		// maintain monotonic decreasing stack
		for len(stack) > 0 && stack[len(stack)-1].Val < num {
			prev = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			// now top element is larger than curr num
			// so num will go to its right
			stack[len(stack)-1].Right = node
		}
		node.Left = prev
		stack = append(stack, node)
	}

	var prev *TreeNode1
	for len(stack) > 0 {
		prev = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return prev
}
