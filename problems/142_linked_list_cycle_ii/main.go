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
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	var meet *ListNode
	slow := head
	fast := head

	for fast != nil {
		if fast.Next == nil {
			return nil
		}

		slow = slow.Next
		fast = fast.Next.Next

		if fast == slow {
			meet = fast

			for meet != head {
				head = head.Next
				meet = meet.Next
			}

			return head
		}
	}

	return nil
}
