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
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			return [][]int{[]int{root.Val}}
		}
	}

	var result [][]int

	leftArrs := pathSum(root.Left, sum-root.Val)
	for _, arr := range leftArrs {
		result = append(result, append([]int{root.Val}, arr...))
	}
	rightArrs := pathSum(root.Right, sum-root.Val)
	for _, arr := range rightArrs {
		result = append(result, append([]int{root.Val}, arr...))
	}

	return result
}
