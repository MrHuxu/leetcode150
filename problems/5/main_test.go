package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_func(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("aba", longestPalindrome("babad"))
	assert.Equal("bb", longestPalindrome("cbbd"))
}
