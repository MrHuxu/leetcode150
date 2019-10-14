package leetcode150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_multiply(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("6", multiply("2", "3"))
	assert.Equal("56088", multiply("123", "456"))
	assert.Equal("1089", multiply("33", "33"))
	assert.Equal("0", multiply("0", "33"))
	assert.Equal("1100", multiply("11", "100"))
}
