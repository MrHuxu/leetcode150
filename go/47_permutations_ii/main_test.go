package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_permuteUnique(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([][]int{
		{1, 1, 2},
		{1, 2, 1},
		{2, 1, 1},
	}, permuteUnique([]int{1, 1, 2}))
}
