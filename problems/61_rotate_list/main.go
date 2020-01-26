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
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	var len int
	tmp := head
	for ; tmp != nil; tmp = tmp.Next {
		len++
	}
	k = k % len

	dummyHead := &ListNode{Next: head}
	fast := dummyHead
	slow := dummyHead
	for ; k > 0; k-- {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	fast.Next = head
	dummyHead.Next = slow.Next
	slow.Next = nil

	return dummyHead.Next
}
