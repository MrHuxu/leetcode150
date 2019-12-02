package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generateMatrix(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([][]int{
		[]int{1, 2, 3},
		[]int{8, 9, 4},
		[]int{7, 6, 5},
	}, generateMatrix(3))
}
