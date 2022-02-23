package main

import (
	"testing"

	. "github.com/MrHuxu/types"

	"github.com/stretchr/testify/assert"
)

func Test_detectCycle(t *testing.T) {
	assert := assert.New(t)

	list := BuildList([]int{1, 2, 3, 4, 5})
	node5 := list.Next.Next.Next.Next
	node3 := list.Next.Next
	node5.Next = node3

	assert.Equal(node3, detectCycle(list))
	assert.Nil(detectCycle(BuildList([]int{1})))
}
