package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_canJump(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(true, canJump([]int{2, 3, 1, 1, 4}))
	assert.Equal(false, canJump([]int{3, 2, 1, 0, 4}))
}
