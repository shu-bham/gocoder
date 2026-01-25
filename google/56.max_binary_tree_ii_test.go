package google

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Tech interview steps:
// 1. Clarifying questions:
//    - What is a maximum binary tree? (A tree where each node's value is greater than any value in its subtree, constructed recursively from an array).
//    - How is the tree constructed? (Root is the max value in the array, left subtree is from the prefix left of max, right subtree from the suffix right of max).
//    - The new value `val` to be inserted corresponds to appending it to the end of the original array. Is that correct? (Yes).
//    - What are the constraints on the values? Are they unique? (Typically, yes).
//    - What should be the output for an empty initial tree (root is nil)? (A new tree with the single node `val`).
//
// 2. High-level approach:
//    - The problem is to insert a value into a max binary tree. This is equivalent to finding the correct position for the new value in the tree structure that would be formed by rebuilding the tree from the original array with the new value appended.
//    - Let the original array be A. The new array is B = A + [val].
//    - Analyze the structure of the new tree based on `val`:
//    - Case 1: If `val` is greater than all elements in A (i.e., `val` > `root.Val`), `val` will be the new root of the tree for B. The entire old tree (from A) becomes the left child of the new root. The right child will be nil. This is because all elements of A are to the left of `val` in B.
//    - Case 2: If `val` is not the greatest element, the old root remains the root of the new tree. Since `val` is appended at the end of A, it falls into the right part of the array relative to the current root. Therefore, we need to recursively insert `val` into the right subtree.
//
// 3. Algorithm:
//    - `insertIntoMaxTree(root, val)`:
//    - Base case: If `root` is nil, it means we've found the spot for insertion. Return a new `TreeNode2` with `val`.
//    - Recursive step 1: If `val > root.Val`, the new value `val` should be the parent of the current subtree. Create a new `TreeNode2` with `val`, set its `Left` child to the current `root`, and return this new node as the new root.
//    - Recursive step 2: If `val < root.Val`, the new value belongs in the right subtree. Recursively call `insertIntoMaxTree` on the right child: `root.Right = insertIntoMaxTree(root.Right, val)`. Return the (original) `root`.
//
// 4. Complexity Analysis:
//    - Time Complexity: O(N) in the worst case. The traversal path is always down the right spine of the tree. In the worst-case scenario (a skewed tree where elements are in decreasing order), the height `h` of the tree is N, so we might traverse all N nodes.
//    - Space Complexity: O(N) in the worst case for the recursion stack, corresponding to the height of the skewed tree.

// Helper function to construct a maximum binary tree from a slice of integers.
// This is needed to create the tree structures for our test cases.
func constructMaximumBinaryTree2(nums []int) *TreeNode2 {
	if len(nums) == 0 {
		return nil
	}
	maxVal := -1
	maxIdx := -1
	for i, v := range nums {
		if v > maxVal {
			maxVal = v
			maxIdx = i
		}
	}
	root := &TreeNode2{Val: maxVal}
	if maxIdx > 0 {
		root.Left = constructMaximumBinaryTree2(nums[:maxIdx])
	}
	if maxIdx < len(nums)-1 {
		root.Right = constructMaximumBinaryTree2(nums[maxIdx+1:])
	}
	return root
}

func TestInsertIntoMaxTree(t *testing.T) {

	t.Run("Insert into an empty tree", func(t *testing.T) {
		var root *TreeNode2 = nil
		val := 5
		expected := &TreeNode2{Val: 5}
		result := insertIntoMaxTree(root, val)
		assert.Equal(t, expected, result)
	})

	t.Run("Insert a new maximum value", func(t *testing.T) {
		// Tree from [4,1,3,2]
		root := constructMaximumBinaryTree2([]int{4, 1, 3, 2})
		val := 5
		// Expected tree from [4,1,3,2,5]
		expected := constructMaximumBinaryTree2([]int{4, 1, 3, 2, 5})
		result := insertIntoMaxTree(root, val)
		assert.Equal(t, expected, result)
	})

	t.Run("Insert a value into the right subtree", func(t *testing.T) {
		// Tree from [5,2,1]
		root := constructMaximumBinaryTree2([]int{5, 2, 1})
		val := 3
		// Expected tree from [5,2,1,3]
		expected := constructMaximumBinaryTree2([]int{5, 2, 1, 3})
		result := insertIntoMaxTree(root, val)
		assert.Equal(t, expected, result)
	})

	t.Run("Insert a value that becomes an intermediate node in the right subtree", func(t *testing.T) {
		// Tree from [10, 8, 12, 4, 2]
		root := constructMaximumBinaryTree2([]int{10, 8, 12, 4, 2})
		val := 6
		// Expected tree from [10, 8, 12, 4, 2, 6]
		expected := constructMaximumBinaryTree2([]int{10, 8, 12, 4, 2, 6})
		result := insertIntoMaxTree(root, val)
		assert.Equal(t, expected, result)
	})

	t.Run("Insert the smallest value as a leaf", func(t *testing.T) {
		// Tree from [4,1,3,2]
		root := constructMaximumBinaryTree2([]int{4, 1, 3, 2})
		val := 0
		// Expected tree from [4,1,3,2,0]
		expected := constructMaximumBinaryTree2([]int{4, 1, 3, 2, 0})
		result := insertIntoMaxTree(root, val)
		assert.Equal(t, expected, result)
	})
}
