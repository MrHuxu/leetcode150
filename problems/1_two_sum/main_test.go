package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_twoSum(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{0, 1}, twoSum([]int{2, 7, 12, 13}, 9))
	assert.Equal([]int{1, 2}, twoSum([]int{3, 2, 4}, 6))
}
