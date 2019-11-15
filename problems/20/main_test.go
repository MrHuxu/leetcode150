package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValid(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(true, isValid("()"))
	assert.Equal(true, isValid("(){}{}"))
	assert.Equal(false, isValid("(]"))
}
