package main

import (
	"testing"

	. "github.com/MrHuxu/types"

	"github.com/stretchr/testify/assert"
)

func Test_inorderTraversal(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		[]int{1, 3, 2},
		inorderTraversal(BuildTree([]interface{}{1, nil, 2, 3})),
	)
}
