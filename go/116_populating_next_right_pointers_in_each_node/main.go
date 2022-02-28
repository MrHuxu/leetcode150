package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// code
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	node := root
	for node != nil {
		iter := node

		node = node.Left
		if node == nil {
			break
		}

		for iter != nil {
			iter.Left.Next = iter.Right
			if iter.Next != nil {
				iter.Right.Next = iter.Next.Left
			}
			iter = iter.Next
		}
	}

	return root
}
