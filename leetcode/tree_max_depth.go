package leetcode

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ld := MaxDepth(root.Left)
	rd := MaxDepth(root.Right)

	return 1 + max(ld, rd)
}
