package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_copyRandomList(t *testing.T) {
	assert := assert.New(t)

	node1 := &Node{Val: 7}
	node2 := &Node{Val: 13}
	node3 := &Node{Val: 11}
	node4 := &Node{Val: 10}
	node5 := &Node{Val: 1}
	node1.Next = node2
	node2.Next = node3
	node2.Random = node1
	node3.Next = node4
	node3.Random = node5
	node4.Next = node5
	node4.Random = node3
	node5.Random = node1

	anotherNode1 := &Node{Val: 7}
	anotherNode2 := &Node{Val: 13}
	anotherNode3 := &Node{Val: 11}
	anotherNode4 := &Node{Val: 10}
	anotherNode5 := &Node{Val: 1}
	anotherNode1.Next = anotherNode2
	anotherNode2.Next = anotherNode3
	anotherNode2.Random = anotherNode1
	anotherNode3.Next = anotherNode4
	anotherNode3.Random = anotherNode5
	anotherNode4.Next = anotherNode5
	anotherNode4.Random = anotherNode3
	anotherNode5.Random = anotherNode1
	assert.Equal(anotherNode1, copyRandomList(node1))
}
