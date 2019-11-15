package main

import (
	"testing"

	. "github.com/MrHuxu/leetcode150/problems/utils"

	"github.com/stretchr/testify/assert"
)

func Test_mergeKLists(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		BuildList([]int{1, 1, 2, 3, 4, 4, 5, 6}),
		mergeKLists(
			[]*ListNode{
				BuildList([]int{1, 4, 5}),
				BuildList([]int{1, 3, 4}),
				BuildList([]int{2, 6}),
			},
		),
	)
}
