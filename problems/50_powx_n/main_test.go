package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_func(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(float64(1024), myPow(2, 10))
	assert.Equal(float64(0.25), myPow(2.00000, -2))
}
