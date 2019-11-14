package leetcode150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_func(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(4, search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	assert.Equal(-1, search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
}
