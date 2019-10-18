package leetcode150

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
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	right := root.Right
	root.Right = root.Left
	root.Left = nil

	for root.Right != nil {
		root = root.Right
	}
	root.Right = right
}
