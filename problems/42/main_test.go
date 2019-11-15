package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_func(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(6, trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	assert.Equal(2, trap([]int{2, 0, 2}))
}
