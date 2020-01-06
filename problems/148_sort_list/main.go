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
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	smaller := &ListNode{}
	smallerPos := smaller
	greater := &ListNode{}
	greaterPos := greater
	equal := &ListNode{Val: head.Val}
	equalPos := equal

	node := head
	for node != nil {
		switch {
		case node.Val < equal.Val:
			smallerPos.Next = node
			smallerPos = smallerPos.Next

		case node.Val > equal.Val:
			greaterPos.Next = node
			greaterPos = greaterPos.Next

		default:
			equalPos.Next = node
			equalPos = equalPos.Next
		}

		node = node.Next
	}
	smallerPos.Next = nil
	greaterPos.Next = nil
	equalPos.Next = nil

	smaller = sortList(smaller.Next)
	greater = sortList(greater.Next)

	dummyHead := &ListNode{Next: smaller}
	tmp := dummyHead
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = equal.Next
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = greater

	return dummyHead.Next
}
