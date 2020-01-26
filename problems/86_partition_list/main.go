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
func partition(head *ListNode, x int) *ListNode {
	smallerHead := &ListNode{Next: head}
	greaterHead := &ListNode{}

	tmp := head
	smaller := smallerHead
	greater := greaterHead
	for tmp != nil {
		if tmp.Val < x {
			smaller.Next = tmp
			smaller = smaller.Next
		} else {
			greater.Next = tmp
			greater = greater.Next
		}
		tmp = tmp.Next
	}
	smaller.Next = greaterHead.Next
	greater.Next = nil

	return smallerHead.Next
}
