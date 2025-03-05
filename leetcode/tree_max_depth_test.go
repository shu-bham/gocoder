package leetcode_test

import (
	leetcode "gocoder/leetcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxDepth(t *testing.T) {
	t.Run("f1 - balanced tree", func(t *testing.T) {
		// Create a balanced tree
		root := &leetcode.TreeNode{
			Val: 1,
			Left: &leetcode.TreeNode{
				Val: 2,
				Left: &leetcode.TreeNode{
					Val: 4,
				},
				Right: &leetcode.TreeNode{
					Val: 5,
				},
			},
			Right: &leetcode.TreeNode{
				Val: 3,
			},
		}
		assert.Equal(t, 3, leetcode.MaxDepth(root))
	})

	t.Run("f2 - left-skewed tree", func(t *testing.T) {
		// Create a left-skewed tree
		root := &leetcode.TreeNode{
			Val: 1,
			Left: &leetcode.TreeNode{
				Val: 2,
				Left: &leetcode.TreeNode{
					Val: 3,
					Left: &leetcode.TreeNode{
						Val: 4,
					},
				},
			},
		}
		assert.Equal(t, 4, leetcode.MaxDepth(root))
	})

	t.Run("f3 - right-skewed tree", func(t *testing.T) {
		// Create a right-skewed tree
		root := &leetcode.TreeNode{
			Val: 1,
			Right: &leetcode.TreeNode{
				Val: 2,
				Right: &leetcode.TreeNode{
					Val: 3,
					Right: &leetcode.TreeNode{
						Val: 4,
					},
				},
			},
		}
		assert.Equal(t, 4, leetcode.MaxDepth(root))
	})

	t.Run("f4 - single node tree", func(t *testing.T) {
		// Create a single node tree
		root := &leetcode.TreeNode{Val: 1}
		assert.Equal(t, 1, leetcode.MaxDepth(root))
	})

	t.Run("f5 - empty tree", func(t *testing.T) {
		// Test empty tree
		var root *leetcode.TreeNode
		assert.Equal(t, 0, leetcode.MaxDepth(root))
	})
}
