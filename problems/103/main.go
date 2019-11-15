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
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var result [][]int
	var fromRight bool

	level := []*TreeNode{root}
	for len(level) != 0 {
		var vals []int
		var nextLevel []*TreeNode

		for i := 0; i < len(level); i++ {
			if fromRight {
				vals = append(vals, level[len(level)-1-i].Val)
			} else {
				vals = append(vals, level[i].Val)
			}

			if level[i].Left != nil {
				nextLevel = append(nextLevel, level[i].Left)
			}
			if level[i].Right != nil {
				nextLevel = append(nextLevel, level[i].Right)
			}
		}

		fromRight = !fromRight
		result = append(result, vals)
		level = nextLevel
	}

	return result
}
