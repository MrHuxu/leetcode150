package main

import (
	"testing"

	. "github.com/MrHuxu/types"
	"github.com/stretchr/testify/assert"
)

func Test_removeNthFromEnd(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		BuildList([]int{1, 2, 3, 5}),
		removeNthFromEnd(BuildList([]int{1, 2, 3, 4, 5}), 2),
	)
}
