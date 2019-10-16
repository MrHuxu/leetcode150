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
func swapPairs(head *ListNode) *ListNode {
	newHead := &ListNode{Next: head}

	tmp := newHead
	for head != nil {
		if head.Next != nil {
			tmp.Next = &ListNode{
				Val:  head.Next.Val,
				Next: &ListNode{Val: head.Val},
			}
			tmp = tmp.Next.Next
			head = head.Next.Next
		} else {
			tmp.Next = &ListNode{Val: head.Val}
			break
		}
	}
	return newHead.Next
}
