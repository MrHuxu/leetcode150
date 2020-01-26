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
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: postorder[len(postorder)-1]}
	var rootIndex int
	for ; rootIndex < len(inorder); rootIndex++ {
		if inorder[rootIndex] == root.Val {
			break
		}
	}
	root.Left = buildTree(inorder[0:rootIndex], postorder[0:rootIndex])
	root.Right = buildTree(inorder[rootIndex+1:len(inorder)], postorder[rootIndex:len(postorder)-1])

	return root
}
