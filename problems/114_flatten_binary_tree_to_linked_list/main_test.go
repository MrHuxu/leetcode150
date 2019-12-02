package main

import (
	"testing"

	. "github.com/MrHuxu/leetcode150/problems/utils"

	"github.com/stretchr/testify/assert"
)

func Test_func(t *testing.T) {
	assert := assert.New(t)

	tree := BuildTree([]interface{}{1, 2, 5, 3, 4, nil, 6})
	flatten(tree)
	assert.Equal(
		BuildTree([]interface{}{1, nil, 2, nil, 3, nil, 4, nil, 5, nil, 6}), tree,
	)
}
