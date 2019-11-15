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
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var result [][]int

	level := []*TreeNode{root}
	for len(level) != 0 {
		var vals []int
		var nextLevel []*TreeNode

		for _, node := range level {
			vals = append(vals, node.Val)

			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}

		result = append([][]int{vals}, result...)
		level = nextLevel
	}

	return result
}
