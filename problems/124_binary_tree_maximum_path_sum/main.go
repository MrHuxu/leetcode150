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
func maxPathSum(root *TreeNode) int {
	result := root.Val

	var traverse func(node *TreeNode) int
	traverse = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftSum := traverse(node.Left)
		rightSum := traverse(node.Right)

		result = max(result, node.Val, node.Val+leftSum+rightSum)

		return max(0, max(leftSum, rightSum)+node.Val)
	}
	traverse(root)

	return result
}

func max(nums ...int) int {
	result := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > result {
			result = nums[i]
		}
	}

	return result
}
