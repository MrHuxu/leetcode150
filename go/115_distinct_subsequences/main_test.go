package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numDistinct(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(1, numDistinct("bag", "bag"))
	assert.Equal(3, numDistinct("rabbbit", "rabbit"))
	assert.Equal(5, numDistinct("babgbag", "bag"))
}
