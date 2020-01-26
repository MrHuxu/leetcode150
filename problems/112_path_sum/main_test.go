package main

import (
	"testing"

	. "github.com/MrHuxu/types"

	"github.com/stretchr/testify/assert"
)

func Test_hasPathSum(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		true, hasPathSum(BuildTree([]interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1}), 22),
	)
}
