package leetcode150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_func(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(2, jump([]int{2, 3, 1, 1, 4}))
}
