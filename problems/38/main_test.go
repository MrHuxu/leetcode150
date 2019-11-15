package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_func(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("1", countAndSay(1))
	assert.Equal("1211", countAndSay(4))
}
