package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_cloneGraph(t *testing.T) {
	assert := assert.New(t)

	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node1.Neighbors = []*Node{node2, node4}
	node2.Neighbors = []*Node{node1, node3}
	node3.Neighbors = []*Node{node2, node4}
	node4.Neighbors = []*Node{node1, node3}

	assert.Equal(cloneGraph(node1), node1)
}
