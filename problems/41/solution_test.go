package leetcode150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_firstMissingPositive(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, firstMissingPositive([]int{1, 2, 0}))
	assert.Equal(2, firstMissingPositive([]int{3, 4, -1, 1}))
	assert.Equal(1, firstMissingPositive([]int{7, 8, 9, 11, 12}))
	assert.Equal(3, firstMissingPositive([]int{0, 2, 2, 1, 1}))
}
