package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_func(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([][]int{
		[]int{-1, -1, 2},
		[]int{0, -1, 1},
	}, threeSum([]int{-1, 0, 1, 2, -1, -4}))

	assert.Equal([][]int{
		[]int{0, 0, 0},
	}, threeSum([]int{0, 0, 0, 0}))
}
