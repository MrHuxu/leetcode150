package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_connect(t *testing.T) {
	assert := assert.New(t)

	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node5 := &Node{Val: 5}
	node7 := &Node{Val: 7}
	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Right = node7

	anotherNode1 := &Node{Val: 1}
	anotherNode2 := &Node{Val: 2}
	anotherNode3 := &Node{Val: 3}
	anotherNode4 := &Node{Val: 4}
	anotherNode5 := &Node{Val: 5}
	anotherNode7 := &Node{Val: 7}
	anotherNode1.Left = anotherNode2
	anotherNode1.Right = anotherNode3
	anotherNode2.Left = anotherNode4
	anotherNode2.Right = anotherNode5
	anotherNode2.Next = anotherNode3
	anotherNode3.Right = anotherNode7
	anotherNode4.Next = anotherNode5
	anotherNode5.Next = anotherNode7
	assert.Equal(anotherNode1, connect(node1))
}
