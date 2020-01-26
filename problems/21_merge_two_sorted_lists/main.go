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
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}

	tmp := head
	for l1 != nil || l2 != nil {
		switch {
		case l1 == nil, (l1 != nil && l2 != nil && l1.Val >= l2.Val):
			tmp.Next = &ListNode{Val: l2.Val}
			l2 = l2.Next

		case l2 == nil, (l1 != nil && l2 != nil && l1.Val < l2.Val):
			tmp.Next = &ListNode{Val: l1.Val}
			l1 = l1.Next
		}

		tmp = tmp.Next
	}

	return head.Next
}
