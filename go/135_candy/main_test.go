package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_candy(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(5, candy([]int{1, 0, 2}))
	assert.Equal(4, candy([]int{1, 2, 2}))
	assert.Equal(7, candy([]int{1, 3, 2, 2, 1}))
	assert.Equal(6, candy([]int{3, 2, 1}))
	assert.Equal(18, candy([]int{1, 6, 10, 8, 7, 3, 2}))
	assert.Equal(11, candy([]int{1, 3, 4, 5, 2}))
}
