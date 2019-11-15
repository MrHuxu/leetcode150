package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_plusOne(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{1, 2, 4}, plusOne([]int{1, 2, 3}))
	assert.Equal([]int{4, 3, 2, 2}, plusOne([]int{4, 3, 2, 1}))
}
