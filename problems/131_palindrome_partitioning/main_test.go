package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_partition(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([][]string{
		[]string{"a", "a", "b"},
		[]string{"aa", "b"},
	}, partition("aab"))
}
