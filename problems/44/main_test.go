package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_func(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(true, isMatch("a", "a*"))
	assert.Equal(false, isMatch("aa", "a"))
	assert.Equal(true, isMatch("aa", "*"))
	assert.Equal(true, isMatch("ca", "?a"))
	assert.Equal(false, isMatch("cb", "?a"))
	assert.Equal(true, isMatch("adceb", "*a*b"))
	assert.Equal(false, isMatch("acdcb", "a*c?b"))
	assert.Equal(false, isMatch("aab", "c*a*b"))
	assert.Equal(false, isMatch("ssi", "ss*?i"))
}
