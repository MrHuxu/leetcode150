package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mySqrt(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(2, mySqrt(4))
	assert.Equal(2, mySqrt(8))
}
