package google

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	leftLca := LowestCommonAncestor(root.Left, p, q)
	rightLca := LowestCommonAncestor(root.Right, p, q)

	if leftLca != nil && rightLca != nil {
		return root
	}

	if leftLca != nil {
		return leftLca
	}

	return rightLca
}
