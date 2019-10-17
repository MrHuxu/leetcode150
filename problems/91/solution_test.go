package leetcode150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numDecodings(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(2, numDecodings("12"))
	assert.Equal(3, numDecodings("226"))
	assert.Equal(0, numDecodings("0"))
	assert.Equal(0, numDecodings("01"))
	assert.Equal(1, numDecodings("10"))
}
