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
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	return generate(1, n)
}

func generate(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	var result []*TreeNode

	for i := start; i <= end; i++ {
		for _, left := range generate(start, i-1) {
			for _, right := range generate(i+1, end) {
				result = append(result, &TreeNode{
					Val:   i,
					Left:  left,
					Right: right,
				})
			}
		}
	}

	return result
}
