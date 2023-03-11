package main

import (
	"testing"

	. "github.com/MrHuxu/types"

	"github.com/stretchr/testify/assert"
)

func Test_sortedListToBST(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		BuildTree([]interface{}{0, -3, 9, -10, nil, 5}),
		sortedListToBST(BuildList([]int{-10, -3, 0, 5, 9})),
	)
}
