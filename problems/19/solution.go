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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tmpHead := &ListNode{Next: head}

	fast := tmpHead
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	slow := tmpHead
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next

	return tmpHead.Next
}
