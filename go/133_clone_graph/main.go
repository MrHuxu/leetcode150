package main

type Node struct {
	Val       int
	Neighbors []*Node
}

// code
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	return helper(node, make(map[int]*Node))
}

func helper(node *Node, nodes map[int]*Node) *Node {
	if node == nil {
		return nil
	}

	nodes[node.Val] = &Node{
		Val: node.Val,
	}
	for _, neighbor := range node.Neighbors {
		if nodes[neighbor.Val] == nil {
			helper(neighbor, nodes)
		}

		nodes[node.Val].Neighbors = append(nodes[node.Val].Neighbors, nodes[neighbor.Val])
	}

	return nodes[node.Val]
}
