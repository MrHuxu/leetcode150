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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	var rootIndex int
	for ; rootIndex < len(inorder); rootIndex++ {
		if inorder[rootIndex] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:1+rootIndex], inorder[0:rootIndex])
	root.Right = buildTree(preorder[rootIndex+1:len(inorder)], inorder[rootIndex+1:len(inorder)])

	return root
}
