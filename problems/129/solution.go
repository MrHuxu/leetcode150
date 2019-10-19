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
func sumNumbers(root *TreeNode) int {
	var sum int

	var traverse func(*TreeNode, int)
	traverse = func(node *TreeNode, pre int) {
		if node == nil {
			return
		}

		curr := pre*10 + node.Val
		switch {
		case node.Left == nil && node.Right == nil:
			sum += curr

		case node.Left == nil:
			traverse(node.Right, curr)

		case node.Right == nil:
			traverse(node.Left, curr)

		default:
			traverse(node.Left, curr)
			traverse(node.Right, curr)

		}
	}
	traverse(root, 0)

	return sum
}
