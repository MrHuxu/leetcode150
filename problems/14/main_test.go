package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestCommonPrefix(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("fl", longestCommonPrefix(
		[]string{"flower", "flow", "flight"},
	))

	assert.Equal("", longestCommonPrefix(
		[]string{"dog", "racecar", "car"},
	))
}
