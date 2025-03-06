package leetcode

/**
 * Definition for a binary tree node.
 */

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ld := MaxDepth(root.Left)
	rd := MaxDepth(root.Right)

	return 1 + max(ld, rd)
}
