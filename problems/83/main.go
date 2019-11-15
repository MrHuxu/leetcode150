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
func deleteDuplicates(head *ListNode) *ListNode {
	newHead := &ListNode{Next: head}

	fast := head
	slow := head
	for fast != nil {
		for fast != nil && fast.Val == slow.Val {
			fast = fast.Next
		}
		slow.Next = fast
		slow = slow.Next
	}

	return newHead.Next
}
