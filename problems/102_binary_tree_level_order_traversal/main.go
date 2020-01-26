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
func levelOrder(root *TreeNode) [][]int {
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

		result = append(result, vals)
		level = nextLevel
	}

	return result
}
