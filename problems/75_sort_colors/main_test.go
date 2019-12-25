package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortColors(t *testing.T) {
	assert := assert.New(t)

	arr := []int{2, 0, 2, 1, 1, 0}
	sortColors(arr)
	assert.Equal([]int{0, 0, 1, 1, 2, 2}, arr)
}
