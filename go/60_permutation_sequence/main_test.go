package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getPermutation(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("213", getPermutation(3, 3))
	assert.Equal("2314", getPermutation(4, 9))
	assert.Equal("1", getPermutation(1, 1))
}
