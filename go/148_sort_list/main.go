package main

import (
	. "github.com/MrHuxu/types"
)

// code
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var smaller *ListNode
	var smallerPos *ListNode
	var greater *ListNode
	var greaterPos *ListNode
	equal := head
	equalPos := equal

	node := head.Next
	for node != nil {
		switch {
		case node.Val < equal.Val:
			if smallerPos == nil {
				smaller = node
				smallerPos = node
			} else {
				smallerPos.Next = node
				smallerPos = smallerPos.Next
			}

		case node.Val > equal.Val:
			if greaterPos == nil {
				greater = node
				greaterPos = node
			} else {
				greaterPos.Next = node
				greaterPos = greaterPos.Next
			}

		default:
			equalPos.Next = node
			equalPos = equalPos.Next
		}

		node = node.Next
	}
	if smallerPos != nil {
		smallerPos.Next = nil
	}
	if greaterPos != nil {
		greaterPos.Next = nil
	}
	equalPos.Next = nil

	smaller = sortList(smaller)
	greater = sortList(greater)

	var result *ListNode
	var resultPos *ListNode
	if smaller != nil {
		result = smaller
		resultPos = result
		for resultPos.Next != nil {
			resultPos = resultPos.Next
		}
		resultPos.Next = equal
	} else {
		result = equal
		resultPos = result
	}
	for resultPos.Next != nil {
		resultPos = resultPos.Next
	}
	resultPos.Next = greater

	return result
}
