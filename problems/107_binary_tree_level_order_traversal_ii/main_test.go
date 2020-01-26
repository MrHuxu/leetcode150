package main

import (
	"testing"

	. "github.com/MrHuxu/types"

	"github.com/stretchr/testify/assert"
)

func Test_levelOrderBottom(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([][]int{
		[]int{15, 7},
		[]int{9, 20},
		[]int{3},
	}, levelOrderBottom(BuildTree([]interface{}{3, 9, 20, nil, nil, 15, 7})))
}
