package leetcode150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_func(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(true, isInterleave("aabcc", "dbbca", "aadbbcbcac"))
	assert.Equal(false, isInterleave("aabcc", "dbbca", "aadbbbaccc"))
}
