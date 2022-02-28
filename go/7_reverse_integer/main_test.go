package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverse(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(321, reverse(123))
	assert.Equal(-321, reverse(-123))
	assert.Equal(21, reverse(120))
	assert.Equal(1, reverse(1))
}
