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
func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	lastHalf := getLastHalf(head)
	reversedLastHalf := reverse(lastHalf)
	merge(head, reversedLastHalf)
}

func getLastHalf(head *ListNode) *ListNode {
	slow := head
	fast := head

	preSlow := &ListNode{Next: slow}

	for fast.Next != nil {
		preSlow = slow
		slow = slow.Next
		if fast.Next.Next != nil {
			fast = fast.Next.Next
		} else {
			break
		}
	}

	preSlow.Next = nil
	return slow
}

func reverse(head *ListNode) *ListNode {
	var result *ListNode
	for head != nil {
		tmp := head
		head = head.Next
		tmp.Next = result
		result = tmp
	}
	return result
}

func merge(head1, head2 *ListNode) *ListNode {
	if head1 == head2 {
		return head1
	}

	dummyHead := &ListNode{}
	dummyHeadPos := dummyHead

	var i int
	for head1 != nil || head2 != nil {
		if i%2 == 0 && head1 != nil {
			dummyHeadPos.Next = head1
			head1 = head1.Next
			dummyHeadPos = dummyHeadPos.Next
		} else {
			dummyHeadPos.Next = head2
			head2 = head2.Next
			dummyHeadPos = dummyHeadPos.Next
		}

		i++
	}

	return dummyHead.Next
}
