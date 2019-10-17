package leetcode150

import (
	"testing"

	. "github.com/MrHuxu/leetcode150/problems/utils"

	"github.com/stretchr/testify/assert"
)

func Test_generateTrees(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]*TreeNode{
		BuildTree([]interface{}{1, nil, 2, nil, 3}),
		BuildTree([]interface{}{1, nil, 3, 2}),
		BuildTree([]interface{}{2, 1, 3}),
		BuildTree([]interface{}{3, 1, nil, nil, 2}),
		BuildTree([]interface{}{3, 2, nil, 1}),
	}, generateTrees(3))

	assert.Equal(([]*TreeNode)(nil), generateTrees(0))
}
