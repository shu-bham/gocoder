package hard

/**
https://leetcode.com/problems/recover-a-tree-from-preorder-traversal/
We run a preorder depth-first search (DFS) on the root of a binary tree.

At each node in this traversal, we output D dashes (where D is the depth of this node), then we output the value of this node.  If the depth of a node is D, the depth of its immediate child is D + 1.  The depth of the root node is 0.

If a node has only one child, that child is guaranteed to be the left child.

Given the output traversal of this traversal, recover the tree and return its root.
*/

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

// time: O(n)
// space: O(n)
func RecoverFromPreorder(traversal string) *TreeNode {
	dashes := 0
	var stack []*TreeNode

	for i := 0; i < len(traversal); {
		// count dashes
		if traversal[i] == '-' {
			dashes++
			i++
		} else {
			// if a no. is encountered
			num := 0
			j := i
			// eval the number
			for j = i; j < len(traversal) && traversal[j] != '-'; j++ {
				num = num*10 + int(traversal[j]-'0')
			}

			node := &TreeNode{Val: num}
			// pop till elements in stack > dashes
			//
			for len(stack) > dashes {
				stack = stack[:len(stack)-1]
			}

			pos := len(stack) - 1
			if pos >= 0 {
				top := stack[pos]
				if top.Left == nil {
					top.Left = node
				} else {
					top.Right = node
				}
			}

			// stack will have all nodes in the order of depth
			stack = append(stack, node)
			// resume from number onwards
			i = j
			// reset
			dashes = 0
		}
	}

	return stack[0]
}
