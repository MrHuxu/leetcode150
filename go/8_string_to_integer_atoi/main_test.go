package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_myAtoi(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(42, myAtoi("42"))
	assert.Equal(-42, myAtoi("    -42"))
	assert.Equal(4193, myAtoi("4193 with words"))
	assert.Equal(0, myAtoi("words and 987"))
	assert.Equal(-2147483648, myAtoi("-91283472332"))
	assert.Equal(2147483647, myAtoi("9223372036854775808"))
}
