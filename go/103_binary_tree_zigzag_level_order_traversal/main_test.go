package main

import (
	"testing"

	. "github.com/MrHuxu/types"

	"github.com/stretchr/testify/assert"
)

func Test_zigzagLevelOrder(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([][]int{
		{3},
		{20, 9},
		{15, 7},
	}, zigzagLevelOrder(BuildTree([]interface{}{3, 9, 20, nil, nil, 15, 7})))
}
