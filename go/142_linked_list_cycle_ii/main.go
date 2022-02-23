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
	if head == nil {
		return nil
	}

	slow := head
	fast := head.Next
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return nil
		}

		fast = fast.Next.Next
		slow = slow.Next
	}

	nodeInList := fast
	listLen := 1
	tmp := nodeInList.Next
	for tmp != nodeInList {
		listLen++
		tmp = tmp.Next
	}

	fast = head
	slow = head
	for i := 0; i < listLen; i++ {
		slow = slow.Next
	}

	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return fast
}
