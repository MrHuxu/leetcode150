package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_subsetsWithDup(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([][]int{
		[]int{},
		[]int{1},
		[]int{1, 2},
		[]int{1, 2, 2},
		[]int{2},
		[]int{2, 2},
	}, subsetsWithDup([]int{1, 2, 2}))
}
