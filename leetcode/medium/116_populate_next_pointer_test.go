package medium_test

import (
	"gocoder/leetcode/medium"
	"testing"
)

func TestConnect(t *testing.T) {

	t.Run("t1", func(t *testing.T) {
		root := createTree()
		medium.Connect(root)
	})
}

func createTree() *medium.Node {

	root := &medium.Node{Val: 1}
	root.Left = &medium.Node{Val: 2}
	root.Right = &medium.Node{Val: 3}
	root.Left.Left = &medium.Node{Val: 4}
	root.Left.Right = &medium.Node{Val: 5}
	root.Right.Left = &medium.Node{Val: 6}
	root.Right.Right = &medium.Node{Val: 7}
	return root
}
