package google

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	//       3
	//      / \
	//     5   1
	//    / \ / \
	//   6  2 0  8
	//     / \
	//    7   4
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 8}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Left.Right.Right = &TreeNode{Val: 4}

	t.Run("p and q in different subtrees", func(t *testing.T) {
		p := root.Left
		q := root.Right
		expected := root
		act := LowestCommonAncestor(root, p, q)
		assert.Equal(t, expected, act)
	})

	t.Run("p is ancestor of q", func(t *testing.T) {
		p := root.Left
		q := root.Left.Right.Right
		expected := root.Left
		act := LowestCommonAncestor(root, p, q)
		assert.Equal(t, expected, act)
	})

	t.Run("q is ancestor of p", func(t *testing.T) {
		p := root.Left.Right.Right
		q := root.Left
		expected := root.Left
		act := LowestCommonAncestor(root, p, q)
		assert.Equal(t, expected, act)
	})

	t.Run("p and q are siblings", func(t *testing.T) {
		p := root.Left.Left
		q := root.Left.Right
		expected := root.Left
		act := LowestCommonAncestor(root, p, q)
		assert.Equal(t, expected, act)
	})

	t.Run("lca is in right subtree", func(t *testing.T) {
		p := root.Right.Left
		q := root.Right.Right
		expected := root.Right
		act := LowestCommonAncestor(root, p, q)
		assert.Equal(t, expected, act)
	})

	t.Run("p and q are deep", func(t *testing.T) {
		p := root.Left.Left
		q := root.Left.Right.Right
		expected := root.Left
		act := LowestCommonAncestor(root, p, q)
		assert.Equal(t, expected, act)
	})
}
