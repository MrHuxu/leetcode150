package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_func(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([][]int{
		[]int{1, 1, 2},
		[]int{1, 2, 1},
		[]int{2, 1, 1},
	}, permuteUnique([]int{1, 1, 2}))
}
