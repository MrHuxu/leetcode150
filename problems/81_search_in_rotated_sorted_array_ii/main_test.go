package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_search(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(true, search([]int{2, 5, 6, 0, 0, 1, 2}, 0))
	assert.Equal(false, search([]int{2, 5, 6, 0, 0, 1, 2}, 3))
}
