package medium

import (
	"math"
)

/**
116. Populating Next Right Pointers in Each Node
https://leetcode.com/problems/populating-next-right-pointers-in-each-node/description/

You are given a perfect binary tree where all leaves are on the same level, and every parent has two children. The binary tree has the following definition:

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.

Initially, all next pointers are set to NULL.

'Amazon', 'Apple', 'Bloomberg', 'Didi', 'Facebook', 'Google', 'Microsoft', 'Nutanix', 'Oracle', 'Sap', 'Vmware'

*/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func Connect(root *Node) *Node {
	if root == nil {
		return root
	}

	if root.Left == nil && root.Right == nil {
		return root
	}

	queue := make([]*Node, 0)
	queue = append(queue, root)

	height := 0
	for len(queue) > 0 {
		// no. of elements at same height
		n := int(math.Pow(2, float64(height)))
		var prev *Node
		for i := 0; i < n; i++ {
			front := queue[0]
			queue = queue[1:]
			if front.Left != nil && front.Right != nil {
				front.Left.Next = front.Right
				queue = append(queue, front.Left)
				queue = append(queue, front.Right)
				if prev != nil {
					prev.Next = front.Left
				}
			}
			prev = front.Right
		}
		height++
	}

	return root
}
