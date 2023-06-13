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
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummyHead := &ListNode{Next: head}
	helper(dummyHead, head, head, nil, k, 0)
	return dummyHead.Next
}

func helper(head, start, curr, reversed *ListNode, k, reversedLen int) {
	switch {
	case reversedLen == k:
		head.Next = reversed
		for i := 0; i < reversedLen; i++ {
			head = head.Next
		}

		helper(head, curr, curr, nil, k, 0)

	case curr == nil:
		head.Next = start

	default:
		helper(
			head, start, curr.Next, &ListNode{Val: curr.Val, Next: reversed}, k, reversedLen+1,
		)
	}
}
