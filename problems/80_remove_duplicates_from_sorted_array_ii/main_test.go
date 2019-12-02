package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeDuplicates(t *testing.T) {
	assert := assert.New(t)

	arr := []int{1, 1, 1, 2, 2, 3}
	assert.Equal(5, removeDuplicates(arr))
	assert.Equal(1, arr[0])
	assert.Equal(1, arr[1])
	assert.Equal(2, arr[2])
	assert.Equal(2, arr[3])
	assert.Equal(3, arr[4])

	arr = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	assert.Equal(7, removeDuplicates(arr))
	assert.Equal(0, arr[0])
	assert.Equal(0, arr[1])
	assert.Equal(1, arr[2])
	assert.Equal(1, arr[3])
	assert.Equal(2, arr[4])
	assert.Equal(3, arr[5])
	assert.Equal(3, arr[6])

	arr = []int{1, 1, 1, 1, 2, 2, 3}
	assert.Equal(5, removeDuplicates((arr)))
	assert.Equal(1, arr[0])
	assert.Equal(1, arr[1])
	assert.Equal(2, arr[2])
	assert.Equal(2, arr[3])
	assert.Equal(3, arr[4])
}
