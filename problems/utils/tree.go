package utils

import "encoding/json"

// TreeNode defines the struct of a node in a numeric binary tree
type TreeNode struct {
	Val   int       `json:"val"`
	Left  *TreeNode `json:"left"`
	Right *TreeNode `json:"right"`
}

// String formats the print content of a node
func (node *TreeNode) String() string {
	bytes, err := json.Marshal(node)

	if err != nil {
		return err.Error()
	}

	return string(bytes)
}

// BuildTree builds a binary tree by the given numeric values
func BuildTree(vals []interface{}) *TreeNode {
	if len(vals) == 0 {
		return nil
	}

	root := &TreeNode{Val: vals[0].(int)}
	vals = vals[1:len(vals)]

	level := []*TreeNode{root}
	for len(vals) != 0 {
		var nextLevel []*TreeNode

		for _, node := range level {
			if node != nil && len(vals) != 0 {
				if vals[0] == nil {
					nextLevel = append(nextLevel, nil)
				} else {
					node.Left = &TreeNode{Val: vals[0].(int)}
					nextLevel = append(nextLevel, node.Left)
				}
				vals = vals[1:len(vals)]

				if len(vals) == 0 {
					break
				} else {
					if vals[0] == nil {
						nextLevel = append(nextLevel, nil)
					} else {
						node.Right = &TreeNode{Val: vals[0].(int)}
						nextLevel = append(nextLevel, node.Right)
					}
					vals = vals[1:len(vals)]
				}
			}
		}

		level = nextLevel
	}

	return root
}
