package leetcode150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_func(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("6", multiply("2", "3"))
	assert.Equal("56088", multiply("123", "456"))
	assert.Equal("3330", singleMulti("111", '3', 1))
	assert.Equal("1240", singleMulti("31", '4', 1))
	assert.Equal("333", add("111", "222"))
	assert.Equal("1110", add("888", "222"))
}
