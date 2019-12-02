package main

import (
	"testing"

	. "github.com/MrHuxu/leetcode150/problems/utils"

	"github.com/stretchr/testify/assert"
)

func Test_func(t *testing.T) {
	assert := assert.New(t)

	tree := BuildTree([]interface{}{1,3,nil,nil,2})
	recoverTree(tree)
	assert.Equal(BuildTree([]interface{}{3, 1, nil, nil, 2}), tree)

	tree = BuildTree([]interface{}{3,1,4,nil,nil,2})
	recoverTree(tree)
	assert.Equal(BuildTree([]interface{}{2,1,4,nil,nil,3}), tree)
}
