package leetcode150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_strStr(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(2, strStr("hello", "ll"))
	assert.Equal(-1, strStr("aaaaa", "bba"))
	assert.Equal(0, strStr("aaaaa", ""))
}
