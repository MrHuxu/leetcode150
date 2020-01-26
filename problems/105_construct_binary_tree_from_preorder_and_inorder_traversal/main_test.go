package main

import (
	"testing"

	. "github.com/MrHuxu/types"

	"github.com/stretchr/testify/assert"
)

func Test_buildTree(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		BuildTree([]interface{}{3, 9, 20, nil, nil, 15, 7}),
		buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}),
	)
}
