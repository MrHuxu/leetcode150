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
func maxPathSum(root *TreeNode) int {
	result := root.Val

	var traverse func(root *TreeNode) int
	traverse = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		leftSum := traverse(root.Left)
		rightSum := traverse(root.Right)

		result = max(result, root.Val, root.Val+leftSum+rightSum)

		return max(0, max(leftSum, rightSum)+root.Val)
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
