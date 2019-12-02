package main

import (
	"testing"

	. "github.com/MrHuxu/leetcode150/problems/utils"

	"github.com/stretchr/testify/assert"
)

func Test_func(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{3, 2, 1}, postorderTraversal(BuildTree([]interface{}{1, nil, 2, 3})))
}
