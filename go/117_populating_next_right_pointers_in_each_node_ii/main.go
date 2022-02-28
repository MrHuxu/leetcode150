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

		var pre *Node
		var nextNode *Node
		for iter != nil {
			if iter.Left != nil {
				if nextNode == nil {
					nextNode = iter.Left
				}

				if pre == nil {
					pre = iter.Left
				} else {
					pre.Next = iter.Left
					pre = pre.Next
				}
			}

			if iter.Right != nil {
				if nextNode == nil {
					nextNode = iter.Right
				}

				if pre == nil {
					pre = iter.Right
				} else {
					pre.Next = iter.Right
					pre = pre.Next
				}
			}

			iter = iter.Next
		}

		node = nextNode
	}

	return root
}
