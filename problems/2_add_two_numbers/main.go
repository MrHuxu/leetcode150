package main

import . "github.com/MrHuxu/types"

// code
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tmp := head

	var plus int
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

		sum += plus
		plus = sum / 10

		tmp.Next = &ListNode{Val: sum % 10}
		tmp = tmp.Next
	}
	if plus == 1 {
		tmp.Next = &ListNode{Val: 1}
	}

	return head.Next
}
