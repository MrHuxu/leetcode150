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
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	newHead := &ListNode{Next: head}
	recursion(newHead, newHead.Next, nil, m, n, 1)

	return newHead.Next
}

func recursion(pre, curr, reversed *ListNode, m, n, seq int) {
	switch {
	case seq < m:
		recursion(curr, curr.Next, reversed, m, n, seq+1)

	case seq >= m && seq < n:
		tmp := curr.Next
		curr.Next = reversed
		recursion(pre, tmp, curr, m, n, seq+1)

	case seq == n:
		pre.Next = curr

		tmp := curr.Next
		curr.Next = reversed
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = tmp
	}
}
