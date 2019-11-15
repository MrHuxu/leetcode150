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
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

	for fast.Next != nil {
		if fast.Next.Next != nil {
			fast = fast.Next.Next
		} else {
			return false
		}

		if fast == slow {
			return true
		}

		slow = slow.Next
	}

	return false
}
