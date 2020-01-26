package main

import (
	"testing"

	. "github.com/MrHuxu/types"

	"github.com/stretchr/testify/assert"
)

func Test_insertionSortList(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		BuildList([]int{1, 2, 3, 4}),
		insertionSortList(BuildList([]int{4, 2, 1, 3})),
	)

	assert.Equal(
		BuildList([]int{-1, 0, 3, 4, 5}),
		insertionSortList(BuildList([]int{-1, 5, 3, 4, 0})),
	)
}
