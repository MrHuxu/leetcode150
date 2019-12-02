package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addBinary(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("100", addBinary("11", "1"))
	assert.Equal("10101", addBinary("1010", "1011"))
}
