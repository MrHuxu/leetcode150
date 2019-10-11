package leetcode150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_func(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, test([]int{1, 2, 0}))
	assert.Equal(2, test([]int{3, 4, -1, 1}))
	assert.Equal(1, test([]int{7, 8, 9, 11, 12}))
}
