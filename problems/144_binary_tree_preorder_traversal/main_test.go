package main

import (
	"testing"

	. "github.com/MrHuxu/types"

	"github.com/stretchr/testify/assert"
)

func Test_preorderTraversal(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{1, 2, 3}, preorderTraversal(BuildTree([]interface{}{1, nil, 2, 3})))
}
