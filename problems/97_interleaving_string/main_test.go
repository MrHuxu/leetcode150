package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isInterleave(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(true, isInterleave("aabcc", "dbbca", "aadbbcbcac"))
	assert.Equal(false, isInterleave("aabcc", "dbbca", "aadbbbaccc"))
}
