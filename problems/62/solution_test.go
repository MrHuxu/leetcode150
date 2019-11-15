package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_uniquePaths(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, uniquePaths(3, 2))
	assert.Equal(28, uniquePaths(7, 3))
}
