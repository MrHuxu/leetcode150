package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxSubArray(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(6, maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
