package main

import (
	"testing"

	. "github.com/MrHuxu/leetcode150/problems/utils"

	"github.com/stretchr/testify/assert"
)

func Test_isValidBST(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		true, isValidBST(BuildTree([]interface{}{2, 1, 3})),
	)
	assert.Equal(
		false, isValidBST(BuildTree([]interface{}{5, 1, 4, nil, nil, 3, 6})),
	)
	assert.Equal(
		false, isValidBST(BuildTree([]interface{}{10, 5, 15, nil, nil, 6, 20})),
	)
}
