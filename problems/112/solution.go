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
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	switch {
	case root.Left == nil && root.Right == nil:
		return sum == root.Val

	case root.Left == nil:
		return hasPathSum(root.Right, sum-root.Val)

	case root.Right == nil:
		return hasPathSum(root.Left, sum-root.Val)

	default:
		return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
	}
}
