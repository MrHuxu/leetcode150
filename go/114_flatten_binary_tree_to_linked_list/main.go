package main

import . "github.com/MrHuxu/types"

// code
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	helper(root)
}

func helper(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Left != nil {
		tail := helper(root.Left)
		tail.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}

	if root.Right != nil {
		return helper(root.Right)
	}

	return root
}
