package main

import (
	"testing"

	. "github.com/MrHuxu/leetcode150/problems/utils"

	"github.com/stretchr/testify/assert"
)

func Test_zigzagLevelOrder(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([][]int{
		[]int{3},
		[]int{20, 9},
		[]int{15, 7},
	}, zigzagLevelOrder(BuildTree([]interface{}{3, 9, 20, nil, nil, 15, 7})))
}
