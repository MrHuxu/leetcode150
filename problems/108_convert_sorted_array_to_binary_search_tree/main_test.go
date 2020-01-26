package main

import (
	"testing"

	. "github.com/MrHuxu/types"

	"github.com/stretchr/testify/assert"
)

func Test_sortedArrayToBST(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		BuildTree([]interface{}{0, -10, 5, nil, -3, nil, 9}),
		sortedArrayToBST([]int{-10, -3, 0, 5, 9}),
	)
}
