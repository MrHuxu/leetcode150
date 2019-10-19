package leetcode150

import (
	. "github.com/MrHuxu/leetcode150/problems/utils"
)

// code
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	newHead := &ListNode{Next: head}

	preNode := head
	node := head.Next
	for node != nil {
		nextPreNode := node
		nextNode := node.Next

		tmp := newHead
		for tmp.Next != node {
			if tmp.Next.Val > node.Val {
				preNode.Next = node.Next
				nextPreNode = preNode

				node.Next = tmp.Next
				tmp.Next = node

				break
			}

			tmp = tmp.Next
		}

		preNode = nextPreNode
		node = nextNode
	}

	return newHead.Next
}
