package leetcode150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_func(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, divide(10, 3))
	assert.Equal(-3, divide(7, -2))
	assert.Equal(2147483647, divide(-2147483648, -1))
	assert.Equal(1073741823, divide(2147483647, 2))
}
