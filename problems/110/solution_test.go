package main

import (
	"testing"

	. "github.com/MrHuxu/leetcode150/problems/utils"

	"github.com/stretchr/testify/assert"
)

func Test_isBalanced(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(true, isBalanced(BuildTree([]interface{}{3, 9, 20, nil, nil, 15, 7})))
	assert.Equal(false, isBalanced(BuildTree([]interface{}{1, 2, 2, 3, 3, nil, nil, 4, 4})))
}
