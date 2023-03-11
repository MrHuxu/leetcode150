package main

import (
	. "github.com/MrHuxu/types"
)

// code
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	var len int
	tmp := head
	for tmp != nil {
		len++
		tmp = tmp.Next
	}
	return helper(head, len)
}

func helper(head *ListNode, len int) *TreeNode {
	if len == 0 || head == nil {
		return nil
	}

	left := helper(head, len/2)
	treeNode := &TreeNode{
		Val:  head.Val,
		Left: left,
	}
	if head.Next != nil {
		*head = *head.Next
		treeNode.Right = helper(head, len-len/2-1)
	}

	return treeNode
}
