package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_largestRectangleArea(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(10, largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	assert.Equal(2, largestRectangleArea([]int{2, 1}))
	assert.Equal(1, largestRectangleArea([]int{1}))
	assert.Equal(9, largestRectangleArea([]int{9, 0}))
	assert.Equal(3, largestRectangleArea([]int{2, 1, 2}))
}
