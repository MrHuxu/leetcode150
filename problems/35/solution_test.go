package leetcode150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_func(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(2, test([]int{1, 3, 5, 6}, 5))
	assert.Equal(1, test([]int{1, 3, 5, 6}, 2))
	assert.Equal(4, test([]int{1, 3, 5, 6}, 7))
	assert.Equal(0, test([]int{1, 3, 5, 6}, 0))
}
