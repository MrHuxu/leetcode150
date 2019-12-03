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
func isValidBST1(root *TreeNode) bool {
	if root == nil {
		return true
	}

	_, _, valid := validate(root)
	return valid
}

func validate(root *TreeNode) (int, int, bool) {
	max := root.Val
	min := root.Val

	if root.Left != nil {
		leftMax, leftMin, leftValid := validate(root.Left)
		if leftValid {
			if leftMax >= root.Val {
				return 0, 0, false
			}
			min = leftMin
		} else {
			return 0, 0, false
		}
	}

	if root.Right != nil {
		rightMax, rightMin, rightValid := validate(root.Right)
		if rightValid {
			if rightMin <= root.Val {
				return 0, 0, false
			}
			max = rightMax
		} else {
			return 0, 0, false
		}
	}

	return max, min, true
}

func isValidBST(root *TreeNode) bool {
	valid := true

	var pre *TreeNode
	var inOrderTraverse func(*TreeNode)
	inOrderTraverse = func(node *TreeNode) {
		if !valid || node == nil {
			return
		}

		inOrderTraverse(node.Left)

		if pre != nil && node.Val <= pre.Val {
			valid = false
		}
		pre = node

		inOrderTraverse(node.Right)
	}
	inOrderTraverse(root)

	return valid
}
