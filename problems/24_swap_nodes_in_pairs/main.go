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
func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}

	tmp := dummyHead
	for head != nil {
		if head.Next != nil {
			tmp.Next = &ListNode{
				Val:  head.Next.Val,
				Next: &ListNode{Val: head.Val},
			}
			tmp = tmp.Next.Next
			head = head.Next.Next
		} else {
			tmp.Next = &ListNode{Val: head.Val}
			break
		}
	}
	return dummyHead.Next
}
