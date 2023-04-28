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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return helper(l1, l2, 0)
}

func helper(l1, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}
	val := carry
	if l1 != nil {
		val += l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		val += l2.Val
		l2 = l2.Next
	}
	return &ListNode{
		Val:  val % 10,
		Next: helper(l1, l2, val/10),
	}
}
