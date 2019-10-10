package leetcode150

type ListNode struct {
	Val  int
	Next *ListNode
}

// code
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	tmpHead := &ListNode{Next: head}
	reverse(tmpHead, head, head, nil, k, 0)
	return tmpHead.Next
}

func reverse(head, start, curr, reversed *ListNode, k, reversedLen int) {
	switch {
	case reversedLen == k:
		head.Next = reversed
		for i := 0; i < reversedLen; i++ {
			head = head.Next
		}

		reverse(head, curr, curr, nil, k, 0)

	case curr == nil:
		head.Next = start

	default:
		reverse(
			head, start, curr.Next, &ListNode{Val: curr.Val, Next: reversed}, k, reversedLen+1,
		)
	}
}
