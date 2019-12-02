package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_func(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([][]int{
		[]int{-2, -1, 1, 2},
		[]int{-2, 0, 0, 2},
		[]int{-1, 0, 0, 1},
	}, fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	assert.Equal([][]int{
		[]int{-3, -2, 2, 3},
		[]int{-3, -1, 1, 3},
		[]int{-3, 0, 0, 3},
		[]int{-3, 0, 1, 2},
		[]int{-2, -1, 0, 3},
		[]int{-2, -1, 1, 2},
		[]int{-2, 0, 0, 2},
		[]int{-1, 0, 0, 1},
	}, fourSum([]int{-3, -2, -1, 0, 0, 1, 2, 3}, 0))
}
