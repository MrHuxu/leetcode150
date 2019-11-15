package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_climbStairs(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(2, climbStairs(2))
	assert.Equal(3, climbStairs(3))
}
