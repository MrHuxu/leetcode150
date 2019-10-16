package leetcode150

import . "github.com/MrHuxu/leetcode150/problems/utils"

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

	newHead := &ListNode{Next: head}
	fast := newHead
	slow := newHead
	for ; k > 0; k-- {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	fast.Next = head
	newHead.Next = slow.Next
	slow.Next = nil

	return newHead.Next
}
