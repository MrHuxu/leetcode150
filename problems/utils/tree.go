package utils

// TreeNode defines the struct of a node in a numeric binary tree
type TreeNode struct {
	Val  int
	Left *TreeNode
	Next *TreeNode
}

// BuildTree builds a binary tree by the given numeric values
func BuildTree(vals []int) *TreeNode {
	head := &TreeNode{}

	tmp := head
	for _, val := range vals {
		tmp.Next = &TreeNode{Val: val}
		tmp = tmp.Next
	}

	return head.Next
}
