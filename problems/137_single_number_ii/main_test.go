package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_func(t *testing.T) {
	assert := assert.New(t)

	// assert.Equal(3, singleNumber([]int{2, 2, 3, 2}))
	// assert.Equal(99, singleNumber([]int{0, 1, 0, 1, 0, 1, 99}))
	assert.Equal(-8, singleNumber([]int{-2, -2, -2, -8}))
	// assert.Equal(-4, singleNumber([]int{-2, -2, 1, 1, -3, 1, -3, -3, -4, -2}))
}
