package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_combinationSum2(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([][]int{
		[]int{1, 1, 6},
		[]int{1, 2, 5},
		[]int{1, 7},
		[]int{2, 6},
	}, combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))

	assert.Equal([][]int{
		[]int{1, 2, 2},
		[]int{5},
	}, combinationSum2([]int{2, 5, 2, 1, 2}, 5))
}
