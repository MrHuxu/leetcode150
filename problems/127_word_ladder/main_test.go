package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ladderLength(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(5, ladderLength(
		"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"},
	))

	assert.Equal(0, ladderLength(
		"hit", "cog", []string{"hot", "dot", "dog", "lot", "log"},
	))
}
