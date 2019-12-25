package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minWindow(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("BANC", minWindow("ADOBECODEBANC", "ABC"))
}
