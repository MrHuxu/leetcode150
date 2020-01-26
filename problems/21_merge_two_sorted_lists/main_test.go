package main

import (
	"testing"

	. "github.com/MrHuxu/types"

	"github.com/stretchr/testify/assert"
)

func Test_mergeTwoLists(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		BuildList([]int{1, 1, 2, 3, 4, 4}),
		mergeTwoLists(
			BuildList([]int{1, 2, 4}),
			BuildList([]int{1, 3, 4}),
		),
	)
}
