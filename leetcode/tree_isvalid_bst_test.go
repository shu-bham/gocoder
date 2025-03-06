package leetcode_test

import (
	"gocoder/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidBST(t *testing.T) {
	t.Run("f1 - valid BST", func(t *testing.T) {
		// Valid BST:       2
		//                /   \
		//               1     3
		root := createTree(intPtr(2), intPtr(1), intPtr(3))
		assert.True(t, leetcode.IsValidBST(root))
	})

	t.Run("f2 - invalid BST due to left child", func(t *testing.T) {
		// Invalid BST:     5
		//                /   \
		//               1     4
		//                    / \
		//                   3   6
		root := createTree(intPtr(5), intPtr(1), intPtr(4), nil, nil, intPtr(3), intPtr(6))
		assert.False(t, leetcode.IsValidBST(root))
	})

	t.Run("f3 - invalid BST due to right child", func(t *testing.T) {
		// Invalid BST:     5
		//                /   \
		//               1     6
		//                    / \
		//                   4   7
		root := createTree(intPtr(5), intPtr(1), intPtr(6), nil, nil, intPtr(4), intPtr(7))
		assert.False(t, leetcode.IsValidBST(root))
	})

	t.Run("f4 - single node tree", func(t *testing.T) {
		// Single node tree
		root := createTree(intPtr(1))
		assert.True(t, leetcode.IsValidBST(root))
	})

	t.Run("f5 - empty tree", func(t *testing.T) {
		// Empty tree
		root := createTree()
		assert.True(t, leetcode.IsValidBST(root))
	})

	t.Run("f6 - complex valid BST", func(t *testing.T) {
		// Valid BST:         8
		//                 /    \
		//                4      10
		//               / \       \
		//              2   6       20
		root := createTree(intPtr(8), intPtr(4), intPtr(10), intPtr(2), intPtr(6), nil, intPtr(20))
		assert.True(t, leetcode.IsValidBST(root))
	})
}

func TestIsValidBSTV2(t *testing.T) {
	t.Run("f1 - valid BST", func(t *testing.T) {
		// Valid BST:       2
		//                /   \
		//               1     3
		root := createTree(intPtr(2), intPtr(1), intPtr(3))
		assert.True(t, leetcode.IsValidBSTV2(root))
	})

	t.Run("f2 - invalid BST due to left child", func(t *testing.T) {
		// Invalid BST:     5
		//                /   \
		//               1     4
		//                    / \
		//                   3   6
		root := createTree(intPtr(5), intPtr(1), intPtr(4), nil, nil, intPtr(3), intPtr(6))
		assert.False(t, leetcode.IsValidBSTV2(root))
	})

	t.Run("f3 - invalid BST due to right child", func(t *testing.T) {
		// Invalid BST:     5
		//                /   \
		//               1     6
		//                    / \
		//                   4   7
		root := createTree(intPtr(5), intPtr(1), intPtr(6), nil, nil, intPtr(4), intPtr(7))
		assert.False(t, leetcode.IsValidBSTV2(root))
	})

	t.Run("f4 - single node tree", func(t *testing.T) {
		// Single node tree
		root := createTree(intPtr(1))
		assert.True(t, leetcode.IsValidBSTV2(root))
	})

	t.Run("f5 - empty tree", func(t *testing.T) {
		// Empty tree
		root := createTree()
		assert.True(t, leetcode.IsValidBSTV2(root))
	})

	t.Run("f6 - complex valid BST", func(t *testing.T) {
		// Valid BST:         8
		//                 /    \
		//                4      10
		//               / \       \
		//              2   6       20
		root := createTree(intPtr(8), intPtr(4), intPtr(10), intPtr(2), intPtr(6), nil, intPtr(20))
		assert.True(t, leetcode.IsValidBSTV2(root))
	})
}

// Helper to create pointers for integers
func intPtr(v int) *int {
	return &v
}

// Helper function to create a binary tree from a level-order slice
func createTree(vals ...*int) *leetcode.TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}

	root := &leetcode.TreeNode{Val: *vals[0]}
	queue := []*leetcode.TreeNode{root}
	i := 1

	for i < len(vals) {
		node := queue[0]
		queue = queue[1:]

		// Left child
		if i < len(vals) && vals[i] != nil {
			node.Left = &leetcode.TreeNode{Val: *vals[i]}
			queue = append(queue, node.Left)
		}
		i++

		// Right child
		if i < len(vals) && vals[i] != nil {
			node.Right = &leetcode.TreeNode{Val: *vals[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}
