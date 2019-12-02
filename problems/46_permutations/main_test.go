package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_permute(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([][]int{
		[]int{1, 2, 3},
		[]int{1, 3, 2},
		[]int{2, 1, 3},
		[]int{2, 3, 1},
		[]int{3, 1, 2},
		[]int{3, 2, 1},
	}, permute([]int{1, 2, 3}))
}
