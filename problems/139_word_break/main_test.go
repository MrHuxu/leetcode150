package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_wordBreak(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(true, wordBreak("leetcode", []string{"leet", "code"}))
	assert.Equal(true, wordBreak("applepenapple", []string{"apple", "pen"}))
	assert.Equal(false, wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}
