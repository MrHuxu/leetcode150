package leetcode150

type ListNode struct {
	Val  int
	Next *ListNode
}

func traverse(n, preLen int, node *ListNode) (totalLen int) {
	if node.Next.Next == nil {
		totalLen = preLen + 1
	} else {
		totalLen = traverse(n, preLen+1, node.Next)
	}

	if totalLen-preLen == n {
		node.Next = node.Next.Next
	}
	return
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tmp := &ListNode{Next: head}
	traverse(n, 0, tmp)
	return tmp.Next
}
