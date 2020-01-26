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
	if head == nil {
		return nil
	}

	dummyHead := &ListNode{Next: head}

	fast := dummyHead.Next
	slow := dummyHead
	for fast != nil && fast.Next != nil {
		if fast.Next.Val == slow.Next.Val {
			for fast.Next != nil && fast.Next.Val == slow.Next.Val {
				fast = fast.Next
			}

			slow.Next = fast.Next
			fast = fast.Next
		} else {
			slow = slow.Next
			fast = fast.Next
		}
	}

	return dummyHead.Next
}
