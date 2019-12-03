package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_func(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]string{
		"cat sand dog",
		"cats and dog",
	}, wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
}
