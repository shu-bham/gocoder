package google

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode2 struct {
	Val   int
	Left  *TreeNode2
	Right *TreeNode2
}

func insertIntoMaxTree(root *TreeNode2, val int) *TreeNode2 {
	if root == nil {
		return &TreeNode2{Val: val}
	}
	if val > root.Val {
		node := &TreeNode2{Val: val}
		node.Left = root
		return node
	}
	root.Right = insertIntoMaxTree(root.Right, val)
	return root
}
