package main

import (
	"testing"

	. "github.com/MrHuxu/types"

	"github.com/stretchr/testify/assert"
)

func Test_swapPairs(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		BuildList([]int{2, 1, 4, 3}),
		swapPairs(
			BuildList([]int{1, 2, 3, 4}),
		),
	)
}
