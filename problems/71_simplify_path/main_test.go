package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_simplifyPath(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("/home", simplifyPath("/home/"))
	assert.Equal("/", simplifyPath("/../"))
	assert.Equal("/home/foo", simplifyPath("/home//foo/"))
	assert.Equal("/c", simplifyPath("/a/./b/../../c/"))
	assert.Equal("/c", simplifyPath("/a/../../b/../c//.//"))
	assert.Equal("/a/b/c", simplifyPath("/a//b////c/d//././/.."))
	assert.Equal("/...", simplifyPath("/..."))
	assert.Equal("/is/here", simplifyPath("/home/of/foo/../../bar/../../is/./here/."))
}
