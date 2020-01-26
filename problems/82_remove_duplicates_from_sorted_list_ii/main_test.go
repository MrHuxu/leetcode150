package main

import (
	"testing"

	. "github.com/MrHuxu/types"

	"github.com/stretchr/testify/assert"
)

func Test_deleteDuplicates(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		BuildList([]int{1, 2, 5}),
		deleteDuplicates(BuildList([]int{1, 2, 3, 3, 4, 4, 5})),
	)

	assert.Equal(
		BuildList([]int{2, 3}),
		deleteDuplicates(BuildList([]int{1, 1, 1, 2, 3})),
	)

	assert.Equal(
		(*ListNode)(nil),
		deleteDuplicates(BuildList([]int{1, 1})),
	)
}
