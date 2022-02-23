package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isNumber(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(isNumber("0"), true)
	assert.Equal(isNumber(" 0.1 "), true)
	assert.Equal(isNumber("abc"), false)
	assert.Equal(isNumber("1 a"), false)
	assert.Equal(isNumber("2e10"), true)
	assert.Equal(isNumber(" -90e3   "), true)
	assert.Equal(isNumber(" 1e"), false)
	assert.Equal(isNumber("e3"), false)
	assert.Equal(isNumber(" 6e-1"), true)
	assert.Equal(isNumber(" 99e2.5 "), false)
	assert.Equal(isNumber("53.5e93"), true)
	assert.Equal(isNumber(" --6 "), false)
	assert.Equal(isNumber("-+3"), false)
	assert.Equal(isNumber("95a54e53"), false)

	assert.Equal(true, isNumber(".1"))
	assert.Equal(true, isNumber("2e0"))
	assert.Equal(false, isNumber("4e+"))
	assert.Equal(true, isNumber("256523.e02"))
	assert.Equal(false, isNumber(" "))
}
