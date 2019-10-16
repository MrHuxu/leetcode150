package leetcode150

import . "github.com/MrHuxu/leetcode150/problems/utils"

// code
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tmp := head

	var pre int
	for l1 != nil || l2 != nil {
		var sum int
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += pre
		pre = sum / 10

		tmp.Next = &ListNode{Val: sum % 10}
		tmp = tmp.Next
	}
	if pre == 1 {
		tmp.Next = &ListNode{Val: 1}
	}

	return head.Next
}
