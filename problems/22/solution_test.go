package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generateParenthesis(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]string{
		"((()))",
		"(()())",
		"(())()",
		"()(())",
		"()()()",
	}, generateParenthesis(3))
}
