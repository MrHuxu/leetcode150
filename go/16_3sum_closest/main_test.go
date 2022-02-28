package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_threeSumClosest(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(2, threeSumClosest([]int{-1, 2, 1, -4}, 1))
	assert.Equal(3, threeSumClosest([]int{0, 1, 2}, 0))
	assert.Equal(3, threeSumClosest([]int{1, 1, 1, 0}, 100))
}
