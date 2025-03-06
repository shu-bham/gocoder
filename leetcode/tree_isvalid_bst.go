package leetcode

import (
	"fmt"
	"math"
)

func IsValidBST(root *TreeNode) bool {
	var res []int
	inorder(root, &res)

	fmt.Println(res)
	for i := 1; i < len(res); i++ {
		if res[i] <= res[i-1] {
			return false
		}
	}
	return true
}

func inorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, res)
	*res = append(*res, root.Val)
	inorder(root.Right, res)
}

// method:2, global max and min value
func IsValidBSTV2(root *TreeNode) bool {
	return isValid(root, math.MinInt, math.MaxInt)
}

func isValid(root *TreeNode, minInt int, maxInt int) bool {
	if root == nil {
		return true
	}
	if minInt < root.Val && root.Val < maxInt {
		return isValid(root.Left, minInt, root.Val) && isValid(root.Right, root.Val, maxInt)
	} else {
		return false
	}
}
