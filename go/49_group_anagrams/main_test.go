package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_groupAnagrams(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([][]string{
		{"eat", "tea", "ate"},
		{"tan", "nat"},
		{"bat"},
	}, groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
