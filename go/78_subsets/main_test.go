package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_subsets(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([][]int{
		[]int{},
		[]int{1},
		[]int{1, 2},
		[]int{1, 2, 3},
		[]int{1, 3},
		[]int{2},
		[]int{2, 3},
		[]int{3},
	}, subsets([]int{1, 2, 3}))
}
