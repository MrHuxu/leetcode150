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
func recoverTree(root *TreeNode) {
	var node1, node2 *TreeNode

	var pre *TreeNode
	var inOrderTraverse func(*TreeNode)
	inOrderTraverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		inOrderTraverse(node.Left)

		if pre != nil && node1 == nil && node.Val < pre.Val {
			node1 = pre
			node2 = node
		}
		if node2 != nil && node.Val < node2.Val {
			node2 = node
		}
		pre = node

		inOrderTraverse(node.Right)
	}
	inOrderTraverse(root)

	tmp := node1.Val
	node1.Val = node2.Val
	node2.Val = tmp
}
