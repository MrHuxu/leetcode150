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
	return helper(nil, head)
}

func helper(pre *ListNode, head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if (pre != nil && pre.Val == head.Val) || (head.Next != nil && head.Next.Val == head.Val) {
		return helper(head, head.Next)
	}

	head.Next = helper(head, head.Next)
	return head
}
