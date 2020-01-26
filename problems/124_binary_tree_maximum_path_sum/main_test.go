package main

import (
	"testing"

	. "github.com/MrHuxu/types"

	"github.com/stretchr/testify/assert"
)

func Test_maxPathSum(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(6, maxPathSum(BuildTree([]interface{}{1, 2, 3})))
	assert.Equal(42, maxPathSum(BuildTree([]interface{}{-10, 9, 20, nil, nil, 15, 7})))
	assert.Equal(-1, maxPathSum(BuildTree([]interface{}{-2, -1})))
}
