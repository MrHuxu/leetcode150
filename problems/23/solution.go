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
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		var next []*ListNode
		for i := 0; i < len(lists); i += 2 {
			if i+1 == len(lists) {
				next = append(next, lists[i])
			} else {
				next = append(next, merge(lists[i], lists[i+1]))
			}
		}
		lists = next
	}

	return lists[0]
}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}

	tmp := head
	for l1 != nil || l2 != nil {
		switch {
		case l1 == nil, (l1 != nil && l2 != nil && l1.Val >= l2.Val):
			tmp.Next = &ListNode{Val: l2.Val}
			l2 = l2.Next

		case l2 == nil, (l1 != nil && l2 != nil && l1.Val < l2.Val):
			tmp.Next = &ListNode{Val: l1.Val}
			l1 = l1.Next
		}

		tmp = tmp.Next
	}

	return head.Next
}
