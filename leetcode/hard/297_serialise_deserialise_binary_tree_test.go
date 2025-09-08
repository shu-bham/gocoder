package hard_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/leetcode/hard"
	"testing"
)

func TestSerDesBinTree(t *testing.T) {

	t.Run("t1", func(t *testing.T) {
		ser := hard.Constructor()
		deser := hard.Constructor()
		tree := createTree()
		data := ser.Serialize(tree)
		ans := deser.Deserialize(data)
		assert.True(t, treesEqual(tree, ans))
	})

	t.Run("empty_tree", func(t *testing.T) {
		ser := hard.Constructor()
		deser := hard.Constructor()
		tree := (*hard.TreeNode)(nil)
		data := ser.Serialize(tree)
		ans := deser.Deserialize(data)
		assert.Nil(t, ans)
	})

	t.Run("single_node", func(t *testing.T) {
		ser := hard.Constructor()
		deser := hard.Constructor()
		tree := &hard.TreeNode{Val: 42}
		data := ser.Serialize(tree)
		ans := deser.Deserialize(data)
		assert.True(t, treesEqual(tree, ans))
	})

	t.Run("left_skewed_tree", func(t *testing.T) {
		ser := hard.Constructor()
		deser := hard.Constructor()
		tree := createLeftSkewedTree()
		data := ser.Serialize(tree)
		ans := deser.Deserialize(data)
		assert.True(t, treesEqual(tree, ans))
	})

	t.Run("right_skewed_tree", func(t *testing.T) {
		ser := hard.Constructor()
		deser := hard.Constructor()
		tree := createRightSkewedTree()
		data := ser.Serialize(tree)
		ans := deser.Deserialize(data)
		assert.True(t, treesEqual(tree, ans))
	})

	t.Run("complete_binary_tree", func(t *testing.T) {
		ser := hard.Constructor()
		deser := hard.Constructor()
		tree := createCompleteBinaryTree()
		data := ser.Serialize(tree)
		ans := deser.Deserialize(data)
		assert.True(t, treesEqual(tree, ans))
	})

	t.Run("failed", func(t *testing.T) {
		ser := hard.Constructor()
		deser := hard.Constructor()
		tree := createBFS3243Tree()
		data := ser.Serialize(tree)
		ans := deser.Deserialize(data)
		assert.True(t, treesEqual(tree, ans))
	})

	t.Run("negative_values", func(t *testing.T) {
		ser := hard.Constructor()
		deser := hard.Constructor()
		tree := createTreeWithNegativeValues()
		data := ser.Serialize(tree)
		ans := deser.Deserialize(data)
		assert.True(t, treesEqual(tree, ans))
	})

	t.Run("large_values", func(t *testing.T) {
		ser := hard.Constructor()
		deser := hard.Constructor()
		tree := createTreeWithLargeValues()
		data := ser.Serialize(tree)
		ans := deser.Deserialize(data)
		assert.True(t, treesEqual(tree, ans))
	})
}

// Helper function to compare two trees for equality
func treesEqual(t1, t2 *hard.TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	if t1.Val != t2.Val {
		return false
	}
	return treesEqual(t1.Left, t2.Left) && treesEqual(t1.Right, t2.Right)
}

// Create the example tree from the problem:
//
//	  1
//	 / \
//	2   3
//	   / \
//	  4   5
func createTree() *hard.TreeNode {
	return &hard.TreeNode{
		Val:  1,
		Left: &hard.TreeNode{Val: 2},
		Right: &hard.TreeNode{
			Val:   3,
			Left:  &hard.TreeNode{Val: 4},
			Right: &hard.TreeNode{Val: 5},
		},
	}
}

// Create a left-skewed tree:
//
//	 1
//	/
//
// 2
// /
// 3
func createLeftSkewedTree() *hard.TreeNode {
	return &hard.TreeNode{
		Val: 1,
		Left: &hard.TreeNode{
			Val:  2,
			Left: &hard.TreeNode{Val: 3},
		},
	}
}

// Create a right-skewed tree:
// 1
//
//	\
//	 2
//	  \
//	   3
func createRightSkewedTree() *hard.TreeNode {
	return &hard.TreeNode{
		Val: 1,
		Right: &hard.TreeNode{
			Val:   2,
			Right: &hard.TreeNode{Val: 3},
		},
	}
}

// Create a complete binary tree:
//
//	     1
//	   /   \
//	  2     3
//	 / \   / \
//	4   5 6   7
func createCompleteBinaryTree() *hard.TreeNode {
	return &hard.TreeNode{
		Val: 1,
		Left: &hard.TreeNode{
			Val:   2,
			Left:  &hard.TreeNode{Val: 4},
			Right: &hard.TreeNode{Val: 5},
		},
		Right: &hard.TreeNode{
			Val:   3,
			Left:  &hard.TreeNode{Val: 6},
			Right: &hard.TreeNode{Val: 7},
		},
	}
}

// Create a tree with negative values:
//
//	   -1
//	  /  \
//	-2    3
//	     / \
//	   -4   5
func createTreeWithNegativeValues() *hard.TreeNode {
	return &hard.TreeNode{
		Val:  -1,
		Left: &hard.TreeNode{Val: -2},
		Right: &hard.TreeNode{
			Val:   3,
			Left:  &hard.TreeNode{Val: -4},
			Right: &hard.TreeNode{Val: 5},
		},
	}
}

// Create a tree with large values:
//
//	     1000
//	    /    \
//	2000      3000
//	         /    \
//	     4000      5000
func createTreeWithLargeValues() *hard.TreeNode {
	return &hard.TreeNode{
		Val:  1000,
		Left: &hard.TreeNode{Val: 2000},
		Right: &hard.TreeNode{
			Val:   3000,
			Left:  &hard.TreeNode{Val: 4000},
			Right: &hard.TreeNode{Val: 5000},
		},
	}
}

// Create a tree with BFS order [3,2,4,3]:
//
//	   3
//	  / \
//	 2   4
//	/
//
// 3
func createBFS3243Tree() *hard.TreeNode {
	return &hard.TreeNode{
		Val: 3,
		Left: &hard.TreeNode{
			Val:  2,
			Left: &hard.TreeNode{Val: 3},
		},
		Right: &hard.TreeNode{Val: 4},
	}
}

// Benchmark tests
func BenchmarkSerialize(b *testing.B) {
	tree := createCompleteBinaryTree()
	codec := hard.Constructor()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		codec.Serialize(tree)
	}
}

func BenchmarkDeserialize(b *testing.B) {
	tree := createCompleteBinaryTree()
	codec := hard.Constructor()
	data := codec.Serialize(tree)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		codec.Deserialize(data)
	}
}

func BenchmarkSerializeDeserialize(b *testing.B) {
	tree := createCompleteBinaryTree()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ser := hard.Constructor()
		deser := hard.Constructor()
		data := ser.Serialize(tree)
		deser.Deserialize(data)
	}
}
