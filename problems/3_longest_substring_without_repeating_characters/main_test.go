package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, lengthOfLongestSubstring("abcabcbb"))
	assert.Equal(1, lengthOfLongestSubstring("bbbbb"))
}
