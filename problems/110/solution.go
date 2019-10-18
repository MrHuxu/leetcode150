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
func isBalanced(root *TreeNode) bool {
	_, balanced := traverse(root)
	return balanced
}

func traverse(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}

	leftDepth, leftBalanced := traverse(node.Left)
	rightDepth, rightBalanced := traverse(node.Right)

	if leftBalanced && rightBalanced {
		if leftDepth > rightDepth {
			return leftDepth + 1, leftDepth-rightDepth <= 1
		}
		return rightDepth + 1, rightDepth-leftDepth <= 1
	}

	return 0, false
}
