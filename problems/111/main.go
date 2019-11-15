package main

import . "github.com/MrHuxu/leetcode150/problems/utils"

// code
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	switch {
	case root.Left == nil && root.Right == nil:
		return 1

	case root.Left == nil:
		depth := minDepth(root.Right)
		return depth + 1

	case root.Right == nil:
		depth := minDepth(root.Left)
		return depth + 1

	default:
		return min(minDepth(root.Left), minDepth(root.Right)) + 1
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
