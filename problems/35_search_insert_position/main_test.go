package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchInsert(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(2, searchInsert([]int{1, 3, 5, 6}, 5))
	assert.Equal(1, searchInsert([]int{1, 3, 5, 6}, 2))
	assert.Equal(4, searchInsert([]int{1, 3, 5, 6}, 7))
	assert.Equal(0, searchInsert([]int{1, 3, 5, 6}, 0))
}
