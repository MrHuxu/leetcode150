package main

import . "github.com/MrHuxu/types"

// code
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast != nil {
		for fast != nil && fast.Val == slow.Val {
			fast = fast.Next
		}
		slow.Next = fast
		slow = slow.Next
	}

	return head
}
