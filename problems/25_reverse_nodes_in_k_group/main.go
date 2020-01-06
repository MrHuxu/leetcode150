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
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummyHead := &ListNode{Next: head}
	reverseKGroupHelper(dummyHead, head, head, nil, k, 0)
	return dummyHead.Next
}

func reverseKGroupHelper(head, start, curr, reversed *ListNode, k, reversedLen int) {
	switch {
	case reversedLen == k:
		head.Next = reversed
		for i := 0; i < reversedLen; i++ {
			head = head.Next
		}

		reverseKGroupHelper(head, curr, curr, nil, k, 0)

	case curr == nil:
		head.Next = start

	default:
		reverseKGroupHelper(
			head, start, curr.Next, &ListNode{Val: curr.Val, Next: reversed}, k, reversedLen+1,
		)
	}
}
