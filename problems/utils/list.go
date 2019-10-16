package utils

import (
	"strconv"
	"strings"
)

// ListNode defines the struct of a node in numeric list
type ListNode struct {
	Val  int
	Next *ListNode
}

// String formats the print content of a node
func (node *ListNode) String() string {
	if node == nil {
		return "nil"
	}

	var vals []string
	for node != nil {
		vals = append(vals, strconv.Itoa(node.Val))
		node = node.Next
	}
	return strings.Join(vals, "->")
}

// BuildList builds a linked list by the given numeric values
func BuildList(vals []int) *ListNode {
	head := &ListNode{}

	tmp := head
	for _, val := range vals {
		tmp.Next = &ListNode{Val: val}
		tmp = tmp.Next
	}

	return head.Next
}
