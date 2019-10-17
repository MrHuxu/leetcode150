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
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return validate(root.Left, root.Right)
}

func validate(a, b *TreeNode) bool {
	switch {
	case a == nil && b == nil:
		return true

	case a == nil, b == nil:
		return false

	default:
		return a.Val == b.Val && validate(a.Left, b.Right) && validate(b.Left, a.Right)
	}
}
