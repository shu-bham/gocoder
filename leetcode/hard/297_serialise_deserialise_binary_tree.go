package hard

import (
	"strconv"
	"strings"
)

/*
*
https://leetcode.com/problems/serialize-and-deserialize-binary-tree/description/
297. Serialize and Deserialize Binary Tree

Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your serialization/deserialization algorithm should work. You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure.

Clarification: The input/output format is the same as how LeetCode serializes a binary tree. You do not necessarily need to follow this format, so please be creative and come up with different approaches yourself.
*/
//type TreeNode struct {
//	Val int
//	Left *TreeNode
//	Right *TreeNode
//}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) Serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	var inorder, preOrder []int
	this.getInorder(root, &inorder)
	this.getPreorder(root, &preOrder)

	inStr := make([]string, len(inorder))
	preStr := make([]string, len(preOrder))

	for i, ele := range inorder {
		inStr[i] = strconv.Itoa(ele)
	}

	for i, ele := range preOrder {
		preStr[i] = strconv.Itoa(ele)
	}

	return strings.Join(preStr, ",") + "|" + strings.Join(inStr, ",")

}

// Deserializes your encoded data to tree.
func (this *Codec) Deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	splits := strings.Split(data, "|")
	if len(splits) != 2 {
		return nil
	}

	var preOrder []int
	preStr := strings.Split(splits[0], ",")
	for _, s := range preStr {
		num, _ := strconv.Atoi(s)
		preOrder = append(preOrder, num)
	}

	var inOrder []int
	inStr := strings.Split(splits[1], ",")
	for _, s := range inStr {
		num, _ := strconv.Atoi(s)
		inOrder = append(inOrder, num)
	}

	return this.buildTree(preOrder, inOrder)

}

func (this *Codec) getInorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	this.getInorder(root.Left, res)
	*res = append(*res, root.Val)
	this.getInorder(root.Right, res)

}

func (this *Codec) getPreorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	*res = append(*res, root.Val)
	this.getPreorder(root.Left, res)
	this.getPreorder(root.Right, res)
}

func (this *Codec) buildTree(preOrder []int, inOrder []int) *TreeNode {
	if len(preOrder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preOrder[0]}

	idx := -1
	for i, ele := range inOrder {
		if ele == root.Val {
			idx = i
			break
		}
	}

	if idx == -1 {
		// shouldm't happen
		return root
	}

	leftInorder := inOrder[:idx]
	rightInorder := inOrder[idx+1:]
	leftPreorder := preOrder[1 : 1+len(leftInorder)]
	rightPreorder := preOrder[1+len(leftInorder):]

	root.Left = this.buildTree(leftPreorder, leftInorder)
	root.Right = this.buildTree(rightPreorder, rightInorder)
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
