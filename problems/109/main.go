package main

import . "github.com/MrHuxu/leetcode150/problems/utils"

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
	if head == nil {
		return nil
	}

	fast := head
	slow := head
	preSlow := &ListNode{Next: slow}
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		preSlow = preSlow.Next
	}

	if fast == slow {
		if slow.Next == nil {
			return &TreeNode{Val: head.Val}
		}
		return &TreeNode{Val: head.Val, Right: &TreeNode{Val: head.Next.Val}}
	}

	root := &TreeNode{Val: slow.Val}

	preSlow.Next = nil
	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(slow.Next)

	return root
}
