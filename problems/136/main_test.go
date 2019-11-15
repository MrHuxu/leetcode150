package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_func(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(1, singleNumber([]int{2, 2, 1}))
	assert.Equal(4, singleNumber([]int{4, 1, 2, 1, 2}))
}
