package hard_test

import (
	"fmt"
	"gocoder/leetcode/hard"
	"testing"
)

func TestRecoverFromPreorder(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		str := "1-2--3---4-5--6---7"
		preorder := hard.RecoverFromPreorder(str)
		fmt.Print("Preorder>>")
		res := ""
		evalPreorder(preorder, &res)
		fmt.Print(res)
	})

	t.Run("2", func(t *testing.T) {
		str := "1-401--349---90--88"
		preorder := hard.RecoverFromPreorder(str)
		fmt.Print("Preorder>>")
		res := ""
		evalPreorder(preorder, &res)
		fmt.Print(res)
	})
}

func evalPreorder(root *hard.TreeNode, s *string) {
	if root == nil {
		*s += "null "
		return
	}
	*s += fmt.Sprintf("%d ", root.Val)
	evalPreorder(root.Left, s)
	evalPreorder(root.Right, s)
}
