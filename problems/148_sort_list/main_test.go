package main

import (
	"testing"

	. "github.com/MrHuxu/leetcode150/problems/utils"

	"github.com/stretchr/testify/assert"
)

func Test_sortList(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		BuildList([]int{1, 2, 3, 4}),
		sortList(BuildList([]int{4, 2, 1, 3})),
	)

	assert.Equal(
		BuildList([]int{-1, 0, 3, 4, 5}),
		sortList(BuildList([]int{-1, 5, 3, 4, 0})),
	)
}
