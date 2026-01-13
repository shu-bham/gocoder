package google

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func LevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	que := make([]*TreeNode, 0)
	que = append(que, root)
	for len(que) > 0 {
		k := len(que)
		temp := make([]int, 0)
		for i := 0; i < k; i++ {
			front := que[0]
			que = que[1:]
			temp = append(temp, (*front).Val)
			if front.Left != nil {
				que = append(que, front.Left)
			}
			if front.Right != nil {
				que = append(que, front.Right)
			}
		}
		res = append(res, temp)
	}
	return res
}
